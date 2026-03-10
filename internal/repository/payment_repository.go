package repository

import (
	"context"
	"database/sql"
	"errors"

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
	query := `INSERT INTO payments (appointment_id, amount, payment_method, status, payment_url) 
	          VALUES ($1, $2, $3, $4, $5) 
	          RETURNING id, created_at, updated_at;`
	err := r.db.QueryRowContext(ctx, query,
		payment.AppointmentID,
		payment.Amount,
		payment.PaymentMethod,
		payment.Status,
		payment.PaymentURL,
	).Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)
	return err
}

// get payment by id
func (r *postgresPaymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	query := `SELECT id, appointment_id, amount, payment_method, status, payment_url, created_at, updated_at 
	          FROM payments WHERE id = $1;`
	row := r.db.QueryRowContext(ctx, query, id)
	var pay domain.Payment
	err := row.Scan(
		&pay.ID,
		&pay.AppointmentID,
		&pay.Amount,
		&pay.PaymentMethod,
		&pay.Status,
		&pay.PaymentURL,
		&pay.CreatedAt,
		&pay.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &pay, nil
}

// get payment by appointment id
func (r *postgresPaymentRepository) GetByAppointmentID(ctx context.Context, appointmentID string) (*domain.Payment, error) {
	query := `SELECT id, appointment_id, amount, payment_method, status, payment_url, created_at, updated_at 
	          FROM payments WHERE appointment_id = $1;`
	row := r.db.QueryRowContext(ctx, query, appointmentID)
	var pay domain.Payment
	err := row.Scan(
		&pay.ID,
		&pay.AppointmentID,
		&pay.Amount,
		&pay.PaymentMethod,
		&pay.Status,
		&pay.PaymentURL,
		&pay.CreatedAt,
		&pay.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &pay, nil
}

// update status payment
func (r *postgresPaymentRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE payments SET status = $1, updated_at = NOW() WHERE id = $2;`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}
