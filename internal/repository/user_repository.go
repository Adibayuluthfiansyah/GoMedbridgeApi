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
