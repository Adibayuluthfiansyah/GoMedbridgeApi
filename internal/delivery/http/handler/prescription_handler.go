package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/middleware"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
)

type PrescriptionHandler struct {
	usecase domain.PrescriptionUseCase
}

func NewPrescriptionHandler(usecase domain.PrescriptionUseCase) *PrescriptionHandler {
	return &PrescriptionHandler{
		usecase: usecase,
	}
}

// create prescription handler
func (h *PrescriptionHandler) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized, user id not found",
		})
		return
	}
	doctorIDStr, ok := userID.(string)
	if !ok {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized, invalid user id format",
		})
		return
	}
	var req domain.PrescriptionCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	prescription, err := h.usecase.CreatePrescription(r.Context(), doctorIDStr, req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusCreated, response.JSONResponse{
		Status:  "success",
		Message: "Prescription created successfully",
		Data:    prescription,
	})
}

// get patient prescription handler
func (h *PrescriptionHandler) GetPatientPrescription(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized, user id not found",
		})
		return
	}
	patientIDStr, ok := userID.(string)
	if !ok {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized, invalid user id format",
		})
		return
	}
	prescriptions, err := h.usecase.GetPatientPrescriptions(r.Context(), patientIDStr)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Prescriptions retrieved successfully",
		Data:    prescriptions,
	})
}
