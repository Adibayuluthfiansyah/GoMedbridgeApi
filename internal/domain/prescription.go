package domain

import (
	"context"
	"time"
)

type Prescription struct {
	ID            string    `json:"id"`
	AppointmentID string    `json:"appointment_id"`
	Medication    string    `json:"medication"`
	Dosage        string    `json:"dosage"`
	Instructions  string    `json:"instructions"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PrescriptionCreateRequest struct {
	AppointmentID string `json:"appointment_id"`
	Medication    string `json:"medication"`
	Dosage        string `json:"dosage"`
	Instructions  string `json:"instructions"`
}

type PrescriptionRepository interface {
	Create(ctx context.Context, prescription *Prescription) error
	GetByAppointmentID(ctx context.Context, appointmentID string) ([]Prescription, error)
	GetByPatientID(ctx context.Context, patientID string) ([]Prescription, error)
}

type PrescriptionUseCase interface {
	CreatePrescription(ctx context.Context, doctorID string, req PrescriptionCreateRequest) (*Prescription, error)
	GetPatientPrescriptions(ctx context.Context, patientID string) ([]Prescription, error)
}
