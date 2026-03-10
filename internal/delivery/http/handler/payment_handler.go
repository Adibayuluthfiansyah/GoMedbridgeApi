package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/middleware"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
)

type PaymentHandler struct {
	usecase domain.PaymentUseCase
}

func NewPaymentHandler(usecase domain.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: usecase,
	}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized from handler",
		})
		return
	}
	patientIDStr, ok := userID.(string)
	if !ok {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized from handler",
		})
		return
	}
	var req domain.PaymentCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	payment, err := h.usecase.CreatePaymentLink(r.Context(), patientIDStr, req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusCreated, response.JSONResponse{
		Status:  "success",
		Message: "Payment created successfully",
		Data:    payment,
	})
}

// handler webhook
func (h *PaymentHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var req domain.PaymentWebhookRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	err = h.usecase.HandleWebhook(r.Context(), req)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Webhook handled successfully",
	})
}
