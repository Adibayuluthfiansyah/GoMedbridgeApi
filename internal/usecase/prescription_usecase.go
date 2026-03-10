package usecase

import (
	"context"
	"errors"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type prescriptionUseCase struct {
	prescriptionRepo domain.PrescriptionRepository
	appointmentRepo  domain.AppointmentRepository
}

func NewPrescriptionUsecase(pRepo domain.PrescriptionRepository, aRepo domain.AppointmentRepository) domain.PrescriptionUseCase {
	return &prescriptionUseCase{
		prescriptionRepo: pRepo,
		appointmentRepo:  aRepo,
	}
}

// create prescription
func (u *prescriptionUseCase) CreatePrescription(ctx context.Context, doctorID string, req domain.PrescriptionCreateRequest) (*domain.Prescription, error) {
	if req.AppointmentID == "" || req.Medication == "" || req.Dosage == "" || req.Instructions == "" {
		return nil, errors.New("All fields are required")
	}
	appointment, err := u.appointmentRepo.GetByID(ctx, req.AppointmentID)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return nil, errors.New("appointment not found")
	}
	if appointment.DoctorID != doctorID {
		return nil, errors.New("unauthorized: you are not assigned to this appointment")
	}
	if appointment.Status != "APPROVED" && appointment.Status != "COMPLETED" {
		return nil, errors.New("cannot prescribe medication for unapproved appointments")
	}
	newPrescription := domain.Prescription{
		AppointmentID: req.AppointmentID,
		Medication:    req.Medication,
		Dosage:        req.Dosage,
		Instructions:  req.Instructions,
	}
	err = u.prescriptionRepo.Create(ctx, &newPrescription)
	if err != nil {
		return nil, err
	}
	return &newPrescription, nil
}

// get patient prescription
func (u *prescriptionUseCase) GetPatientPrescriptions(ctx context.Context, patientID string) ([]domain.Prescription, error) {
	prescription, err := u.prescriptionRepo.GetByPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}
	if prescription == nil {
		return []domain.Prescription{}, nil
	}
	return prescription, nil
}
