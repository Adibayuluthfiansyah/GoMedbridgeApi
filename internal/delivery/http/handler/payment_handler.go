package handler

import "github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"

type PaymentHandler struct {
	usecase domain.PaymentUseCase
}

func NewPaymentHandler(usecase domain.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: usecase,
	}
}
