package domain

import (
	"context"
	"time"
)

type Payment struct {
	ID            string    `json:"id"`
	AppointmentID string    `json:"appointment_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	PaymentURL    string    `json:"payment_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentCreateRequest struct {
	AppointmentID string  `json:"appointment_id"`
	Amount        float64 `json:"amount"`
}

type PaymentWebhookRequest struct {
	PaymentID     string `json:"payment_id"`
	PaymentStatus string `json:"payment_status"`
}

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	GetByID(ctx context.Context, id string) (*Payment, error)
	GetByAppointmentID(ctx context.Context, appointmentID string) (*Payment, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}

type PaymentUseCase interface {
	CreatePaymentLink(ctx context.Context, patientID string, req PaymentCreateRequest) (*Payment, error)
	HandleWebhook(ctx context.Context, req PaymentWebhookRequest) error
}
