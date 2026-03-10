package repository

import (
	"context"
	"database/sql"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type postgresPrescriptionRepository struct {
	db *sql.DB
}

func NewPostgresPrescriptionRepository(db *sql.DB) domain.PrescriptionRepository {
	return &postgresPrescriptionRepository{
		db: db,
	}
}

// create prescription
func (r *postgresPrescriptionRepository) Create(ctx context.Context, prescription *domain.Prescription) error {
	query := `INSERT INTO prescriptions (appointment_id, medication, dosage, instructions) VALUES ($1 ,$2, $3, $4) RETURNING id,created_at,updated_at;`
	err := r.db.QueryRowContext(
		ctx, query, prescription.AppointmentID,
		prescription.Medication,
		prescription.Dosage,
		prescription.Instructions,
	).Scan(&prescription.ID, &prescription.CreatedAt, &prescription.UpdatedAt)
	return err
}

// get appointment by id
func (r *postgresPrescriptionRepository) GetByAppointmentID(ctx context.Context, appointmentID string) ([]domain.Prescription, error) {
	query := `SELECT id, appointment_id, medication, dosage, instructions, created_at, updated_at FROM prescriptions WHERE appointment_id = $1 ORDER BY created_at DESC;`
	rows, err := r.db.QueryContext(ctx, query, appointmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prescriptions []domain.Prescription
	for rows.Next() {
		var pre domain.Prescription
		err := rows.Scan(&pre.ID, &pre.AppointmentID, &pre.Medication, &pre.Dosage, &pre.Instructions, &pre.CreatedAt, &pre.UpdatedAt)
		if err != nil {
			return nil, err
		}
		prescriptions = append(prescriptions, pre)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return prescriptions, nil
}

// get patient by id
func (r *postgresPrescriptionRepository) GetByPatientID(ctx context.Context, patientID string) ([]domain.Prescription, error) {
	query := `SELECT pre.id, pre.appointment_id, pre.medication, pre.dosage, pre.instructions, pre.created_at, pre.updated_at FROM prescriptions pre JOIN appointments app ON pre.appointment_id = app.id WHERE app.patient_id = $1 ORDER BY pre.created_at DESC;`
	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prescription []domain.Prescription
	for rows.Next() {
		var pre domain.Prescription
		err := rows.Scan(&pre.ID, &pre.AppointmentID, &pre.Medication, &pre.Dosage, &pre.Instructions, &pre.CreatedAt, &pre.UpdatedAt)
		if err != nil {
			return nil, err
		}
		prescription = append(prescription, pre)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return prescription, nil
}
