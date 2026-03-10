package usecase

import (
	"context"

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
	return nil, nil
}

// get patient prescription
func (u *prescriptionUseCase) GetPatientPrescriptions(ctx context.Context, patientID string) ([]domain.Prescription, error) {
	return nil, nil
}
