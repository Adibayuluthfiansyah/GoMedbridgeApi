package usecase

import (
	"context"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type PaymentUseCase struct {
	repo domain.PaymentRepository
}

func NewPaymentUsecase(repo domain.PaymentRepository) domain.PaymentUseCase {
	return &PaymentUseCase{
		repo: repo,
	}
}

// create payment
func (u *PaymentUseCase) CreatePaymentLink(ctx context.Context, patientID string, req domain.PaymentCreateRequest) (*domain.Payment, error) {
	return nil, nil
}

// handle webhook
func (u *PaymentUseCase) HandleWebhook(ctx context.Context, req domain.PaymentWebhookRequest) error {
	return nil
}
