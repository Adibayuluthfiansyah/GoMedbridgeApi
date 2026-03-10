package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type PaymentUseCase struct {
	paymentRepo     domain.PaymentRepository
	appointmentRepo domain.AppointmentRepository
}

func NewPaymentUsecase(pRepo domain.PaymentRepository, aRepo domain.AppointmentRepository) domain.PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo:     pRepo,
		appointmentRepo: aRepo,
	}
}

// create payment
func (u *PaymentUseCase) CreatePaymentLink(ctx context.Context, patientID string, req domain.PaymentCreateRequest) (*domain.Payment, error) {
	if req.AppointmentID == "" || req.Amount <= 0 {
		return nil, errors.New("appointmentid and valid amount are required")
	}
	appointment, err := u.appointmentRepo.GetByID(ctx, req.AppointmentID)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return nil, errors.New("appointment not found")
	}
	if appointment.PatientID != patientID {
		return nil, errors.New("unauthorized: you are not assigned to this appointment")
	}
	if appointment.Status != "APPROVED" {
		return nil, errors.New("cannot create payment for unapproved appointments")
	}
	existingPayment, err := u.paymentRepo.GetByAppointmentID(ctx, req.AppointmentID)
	if err != nil {
		return nil, err
	}
	if existingPayment != nil {
		return nil, errors.New("payment already exists for this appointment")
	}
	mockPaymentURL := fmt.Sprintf("https://app.sandbox.midtrans.com/payment/%s-mock-link", req.AppointmentID[:8])
	newPayment := domain.Payment{
		AppointmentID: req.AppointmentID,
		Amount:        req.Amount,
		PaymentMethod: "BANK_TRANSFER",
		Status:        "PENDING",
		PaymentURL:    mockPaymentURL,
	}
	err = u.paymentRepo.Create(ctx, &newPayment)
	if err != nil {
		return nil, err
	}

	return &newPayment, nil
}

// handle webhook
func (u *PaymentUseCase) HandleWebhook(ctx context.Context, req domain.PaymentWebhookRequest) error {
	if req.PaymentID == "" || req.PaymentStatus == "" {
		return errors.New("paymentid and paymentstatus are required")
	}
	payment, err := u.paymentRepo.GetByID(ctx, req.PaymentID)
	if err != nil {
		return err
	}
	if payment == nil {
		return errors.New("payment not found")
	}
	var finalStatus string
	switch req.PaymentStatus {
	case "settlement", "capture", "PAID":
		finalStatus = "PAID"
	case "canccel", "expire", "deny":
		finalStatus = "FAILED"
	default:
		finalStatus = "PENDING"
	}
	err = u.paymentRepo.UpdateStatus(ctx, req.PaymentID, finalStatus)
	if err != nil {
		return err
	}
	return nil
}
