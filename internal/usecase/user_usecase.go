package usecase

import (
	"context"
	"errors"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo      domain.UserRepository
	jwtSecret string
}

func NewUserUsecase(repo domain.UserRepository, jwtSecret string) domain.UserUseCase {
	return &userUseCase{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

// register
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

// login
func (u *userUseCase) Login(ctx context.Context, req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	user, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.UserLoginResponse{}, err
	}
	if user == nil {
		return domain.UserLoginResponse{}, errors.New("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return domain.UserLoginResponse{}, errors.New("Invalid Email and Password")
	}
	tokenString, err := jwt.GenerateToken(user.ID, user.Role, u.jwtSecret)
	if err != nil {
		return domain.UserLoginResponse{}, errors.New("Failed to generate jwt token")
	}
	return domain.UserLoginResponse{
		Token: tokenString,
	}, nil
}

// update profile
func (u *userUseCase) UpdateProfile(ctx context.Context, userID string, req domain.UserUpdateRequest) error {
	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.Name = req.Name
	err = u.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// get by id
func (u *userUseCase) GetByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// get list doctor
func (u *userUseCase) GetDoctors(ctx context.Context) ([]domain.User, error) {
	doctors, err := u.repo.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}
	if doctors == nil {
		return []domain.User{}, nil
	}
	return doctors, nil
}
