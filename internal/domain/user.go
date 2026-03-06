package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	GetDoctors(ctx context.Context) ([]User, error)
}

type UserUseCase interface {
	Register(ctx context.Context, req UserRegisterRequest) error
	Login(ctx context.Context, req UserLoginRequest) (UserLoginResponse, error)
	UpdateProfile(ctx context.Context, userID string, req UserUpdateRequest) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetDoctors(ctx context.Context) ([]User, error)
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
