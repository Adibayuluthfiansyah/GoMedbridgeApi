package domain

import (
	"context"
	"time"
)

type Appointment struct {
	ID              string    `json:"id"`
	PatientID       string    `json:"patient_id"`
	DoctorID        string    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type AppointmentCreateRequest struct {
	DoctorID        string `json:"doctor_id"`
	AppointmentDate string `json:"appointment_date"`
}

type AppointmentUpdateStatusRequest struct {
	Status string `json:"status"`
}

type AppointmentRepository interface {
	Create(ctx context.Context, appointment *Appointment) error
	GetByID(ctx context.Context, id string) (*Appointment, error)
	UpdateStatus(ctx context.Context, id string, req AppointmentUpdateStatusRequest) error
	GetPatientID(ctx context.Context, patientID string) ([]Appointment, error)
}

type AppointmentUseCase interface {
	CreateAppointment(ctx context.Context, patientID string, req AppointmentCreateRequest) (*Appointment, error)
	UpdateAppointmentStatus(ctx context.Context, doctorID string, appointmentID string, req AppointmentUpdateStatusRequest) error
	GetByID(ctx context.Context, id string) (*Appointment, error)
	GetPatientAppointments(ctx context.Context, patientID string) ([]Appointment, error)
}
