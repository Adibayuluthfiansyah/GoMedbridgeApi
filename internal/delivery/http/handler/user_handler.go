package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func NewUserHandler(useCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: useCase,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	err = h.usecase.Register(r.Context(), req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusCreated, response.JSONResponse{
		Status:  "Succes",
		Message: "User Success Registered",
	})
}
