package usecase

import (
	"context"
	"errors"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) Register(ctx context.Context, req domain.UserRegisterRequest) error {
	existingUser, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("Email Already Registered")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}
	err = u.repo.Create(ctx, &newUser)
	if err != nil {
		return err
	}
	return nil
}
