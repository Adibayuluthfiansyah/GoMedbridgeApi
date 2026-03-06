package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

// create user
func (r *postgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password_hash, role) VALUES ($1 , $2, $3, $4) RETURNING id;`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.PasswordHash, user.Role).Scan(&user.ID)
	return err
}

func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password_hash, role, created_at, updated_at FROM users WHERE email = $1;`
	row := r.db.QueryRowContext(ctx, query, email)
	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// get by id
func (r *postgresUserRepository) GetByID(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password_hash, role, created_at, updated_at FROM users WHERE id = $1;`
	row := r.db.QueryRowContext(ctx, query, email)
	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// update user
func (r *postgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, updated_at = NOW() WHERE id = $2;`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.ID)
	return err
}

// get list doctor
func (r *postgresUserRepository) GetDoctors(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, name, email, role, created_at, updated_at FROM users WHERE role = 'doctor';`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var doctors []domain.User
	for rows.Next() {
		var doctor domain.User
		err := rows.Scan(
			&doctor.ID, &doctor.Name, &doctor.Email,
			&doctor.Role, &doctor.CreatedAt, &doctor.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}
