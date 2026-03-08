package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type postgresAppointmentRepository struct {
	db *sql.DB
}

func NewPostgresAppointmentRepository(db *sql.DB) domain.AppointmentRepository {
	return &postgresAppointmentRepository{
		db: db,
	}
}

// create appointment
func (r *postgresAppointmentRepository) Create(ctx context.Context, appointment *domain.Appointment) error {
	query := `INSERT INTO appointments (patient_id, doctor_id, appointment_date, status) VALUES ($1, $2, $3, $4) RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, appointment.PatientID, appointment.DoctorID, appointment.AppointmentDate, appointment.Status).Scan(&appointment.ID)
	return err
}

// update status appointment
func (r *postgresAppointmentRepository) UpdateStatus(ctx context.Context, id string, appointment domain.AppointmentUpdateStatusRequest) error {
	query := `UPDATE appointments SET status = $1, updated_at = NOW() WHERE id =$2;`
	_, err := r.db.ExecContext(ctx, query, appointment.Status, id)
	return err
}

// get by id
func (r *postgresAppointmentRepository) GetByID(ctx context.Context, id string) (*domain.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, appointment_date, status, created_at, updated_at FROM appointments WHERE id = $1;`
	row := r.db.QueryRowContext(ctx, query, id)
	var appointment domain.Appointment
	err := row.Scan(
		&appointment.ID,
		&appointment.PatientID,
		&appointment.DoctorID,
		&appointment.AppointmentDate,
		&appointment.Status,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &appointment, nil
}

// get patient appointment
func (r *postgresAppointmentRepository) GetPatientID(ctx context.Context, patientID string) ([]domain.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, appointment_date, status, created_at, updated_at FROM appointments WHERE patiend_ID = $1 ORDER BY appointment_date DESC;`
	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var appointments []domain.Appointment
	for rows.Next() {
		var app domain.Appointment
		err := rows.Scan(
			&app.ID,
			&app.PatientID,
			&app.DoctorID,
			&app.AppointmentDate,
			&app.Status,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, app)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}
