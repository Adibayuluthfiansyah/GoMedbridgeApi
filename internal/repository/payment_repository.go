package repository

import (
	"context"
	"database/sql"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type postgresPaymentRepository struct {
	db *sql.DB
}

func NewPostgresPaymentRepository(db *sql.DB) domain.PaymentRepository {
	return &postgresPaymentRepository{
		db: db,
	}
}

// create payment
func (r *postgresPaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	return nil
}

// get payment by id
func (r *postgresPaymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	return nil, nil
}

// get payment by appointment id
func (r *postgresPaymentRepository) GetByAppointmentID(ctx context.Context, appointmentID string) (*domain.Payment, error) {
	return nil, nil
}

// update status payment
func (r *postgresPaymentRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	return nil
}
