package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/middleware"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
)

type AppointmentHandler struct {
	usecase domain.AppointmentUseCase
}

func NewAppointmentHandler(useCase domain.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{
		usecase: useCase,
	}
}

// create handler
func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unathorized : patient ID not found",
		})
		return
	}
	patientIDStr := userID.(string)
	var req domain.AppointmentCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	appointment, err := h.usecase.CreateAppointment(r.Context(), patientIDStr, req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusCreated, response.JSONResponse{
		Status:  "success",
		Message: "Appointment created successfully",
		Data:    appointment,
	})
}

// update status handler
func (h *AppointmentHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	doctorID := r.Context().Value(middleware.UserIDKey)
	if doctorID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unathorized : doctor ID not found",
		})
		return
	}
	doctorIDStr, ok := doctorID.(string)
	if !ok {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unathorized : doctor ID not found or wrong format",
		})
		return
	}

	appointmentID := r.PathValue("id")
	if appointmentID == "" {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Appointment ID is required in URL",
		})
		return
	}

	var req domain.AppointmentUpdateStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	err = h.usecase.UpdateAppointmentStatus(r.Context(), doctorIDStr, appointmentID, req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Appointment status updated successfully",
	})
}

// get by id appointments
func (h *AppointmentHandler) GetAppointmentsByID(w http.ResponseWriter, r *http.Request) {
	appointmentID := r.PathValue("id")
	if appointmentID == "" {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Appointments ID required",
		})
		return
	}
	appointment, err := h.usecase.GetByID(r.Context(), appointmentID)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Appointments retrieved successfully",
		Data:    appointment,
	})
}
