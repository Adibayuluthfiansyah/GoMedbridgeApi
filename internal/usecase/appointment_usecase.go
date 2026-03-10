package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type appointmentUseCase struct {
	repo domain.AppointmentRepository
}

func NewAppointmentUsecase(repo domain.AppointmentRepository) domain.AppointmentUseCase {
	return &appointmentUseCase{
		repo: repo,
	}
}

// create appointement
func (u *appointmentUseCase) CreateAppointment(ctx context.Context, patientID string, req domain.AppointmentCreateRequest) (*domain.Appointment, error) {
	parsedDate, err := time.Parse(time.RFC3339, req.AppointmentDate)
	if err != nil {
		return nil, errors.New("Invalid appointment format date")
	}
	newAppointment := domain.Appointment{
		PatientID:       patientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: parsedDate,
		Status:          "PENDING",
	}
	err = u.repo.Create(ctx, &newAppointment)
	if err != nil {
		return nil, err
	}
	return &newAppointment, nil
}

// update status appointment
func (u *appointmentUseCase) UpdateAppointmentStatus(ctx context.Context, doctorID string, appointmentdID string, req domain.AppointmentUpdateStatusRequest) error {
	appointment, err := u.repo.GetByID(ctx, appointmentdID)
	if err != nil {
		return err
	}
	if appointment == nil {
		return errors.New("appointment not found")
	}
	if appointment.DoctorID != doctorID {
		return errors.New("unauthorized: you are not assigned to this appointment")
	}
	if req.Status != "APPROVED" && req.Status != "REJECTED" {
		return errors.New("invalid status: status must be APPROVED or REJECTED")
	}
	err = u.repo.UpdateStatus(ctx, appointmentdID, req)
	if err != nil {
		return err
	}
	return nil
}

// get by id
func (u *appointmentUseCase) GetByID(ctx context.Context, id string) (*domain.Appointment, error) {
	appointment, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return nil, errors.New("appointment not found")
	}
	return appointment, nil
}

// get appointment list
func (u *appointmentUseCase) GetPatientAppointments(ctx context.Context, patientID string) ([]domain.Appointment, error) {
	appointment, err := u.repo.GetPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return []domain.Appointment{}, nil
	}
	return appointment, nil
}

// get doctor appointment
func (u *appointmentUseCase) GetDoctorAppointments(ctx context.Context, doctorID string) ([]domain.Appointment, error) {
	appointment, err := u.repo.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return nil, err
	}
	if appointment == nil {
		return []domain.Appointment{}, nil
	}
	return appointment, nil
}
