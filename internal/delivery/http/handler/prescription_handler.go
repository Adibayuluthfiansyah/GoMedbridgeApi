package handler

import (
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
)

type PrescriptionHandler struct {
	usecase domain.PrescriptionUseCase
}

func NewPrescriptionHandler(usecase domain.PrescriptionUseCase) *PrescriptionHandler {
	return &PrescriptionHandler{
		usecase: usecase,
	}
}

//create prescription handler

//get patient prescription handler

//get appointment prescription handler
