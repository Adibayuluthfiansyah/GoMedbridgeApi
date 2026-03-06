package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/domain"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/middleware"
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

// regis handler
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

// login handler
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	res, err := h.usecase.Login(r.Context(), req)
	if err != nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "succes",
		Message: "Login Successful",
		Data:    res,
	})
}

// get profile
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized from handler",
		})
		return
	}
	userIDStr := userID.(string)

	user, err := h.usecase.GetByID(r.Context(), userIDStr)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Welcome to your profile!",
		Data:    user,
	})
}

// update profile
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}
	userIDStr := userID.(string)
	var req domain.UserUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: "Invalid Request Body",
		})
		return
	}
	err = h.usecase.UpdateProfile(r.Context(), userIDStr, req)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "success",
		Message: "Update Profile Successfully",
	})
}

// get doctors
func (h *UserHandler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	doctors, err := h.usecase.GetDoctors(r.Context())
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.JSONResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	response.WriteJSON(w, http.StatusOK, response.JSONResponse{
		Status:  "Success",
		Message: "List of doctors retrieved successfully",
		Data:    doctors,
	})

}
