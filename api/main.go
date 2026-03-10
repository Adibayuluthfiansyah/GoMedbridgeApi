package main

import (
	"log"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/delivery/http/handler"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/config"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/database"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/repository"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/usecase"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// database.RunMigrations(db)

	mux := http.NewServeMux()

	// repo
	userRepo := repository.NewPostgresUserRepository(db)
	appointmentRepo := repository.NewPostgresAppointmentRepository(db)
	prescriptionRepo := repository.NewPostgresPrescriptionRepository(db)
	paymentRepo := repository.NewPostgresPaymentRepository(db)

	jwtSecret := cfg.JWTSecret

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepo, jwtSecret)
	appointmentUseCase := usecase.NewAppointmentUsecase(appointmentRepo)
	prescriptionUsecase := usecase.NewPrescriptionUsecase(prescriptionRepo, appointmentRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, appointmentRepo)

	// handler
	userHandler := handler.NewUserHandler(userUsecase)
	appointmentHandler := handler.NewAppointmentHandler(appointmentUseCase)
	prescriptionHandler := handler.NewPrescriptionHandler(prescriptionUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)

	// private route endpoints
	updateProfileEndpoint := http.HandlerFunc(userHandler.UpdateProfile)
	profileEndpoint := http.HandlerFunc(userHandler.GetProfile)
	appointmendEndpoint := http.HandlerFunc(appointmentHandler.CreateAppointment)
	updateAppointmentStatusEndpoint := http.HandlerFunc(appointmentHandler.UpdateStatus)
	getAppointmentEndpoint := http.HandlerFunc(appointmentHandler.GetAppointmentsByID)
	getPatientAppointmentEndpoint := http.HandlerFunc(appointmentHandler.GetPatientAppointment)
	getDoctorAppointmentEndpoint := http.HandlerFunc(appointmentHandler.GetDoctorAppointments)
	createPrescriptionEndpoint := http.HandlerFunc(prescriptionHandler.CreatePrescription)
	getPatientPrescriptionEndpoint := http.HandlerFunc(prescriptionHandler.GetPatientPrescription)
	createPaymentEndpoint := http.HandlerFunc(paymentHandler.CreatePayment)

	//route public
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("GET /doctors", userHandler.GetDoctors)
	mux.HandleFunc("POST /payments/webhook", paymentHandler.HandleWebhook)

	//route private
	mux.Handle("PUT /profile", middleware.Auth(jwtSecret)(updateProfileEndpoint))
	mux.Handle("GET /profile", middleware.Auth(jwtSecret)(profileEndpoint))
	mux.Handle("POST /appointments", middleware.Auth(jwtSecret)(appointmendEndpoint))
	mux.Handle("PUT /appointments/{id}/status", middleware.Auth(jwtSecret)(updateAppointmentStatusEndpoint))
	mux.Handle("GET /appointments/{id}", middleware.Auth(jwtSecret)(getAppointmentEndpoint))
	mux.Handle("GET /appointments", middleware.Auth(jwtSecret)(getPatientAppointmentEndpoint))
	mux.Handle("GET /appointments/doctor", middleware.Auth(jwtSecret)(getDoctorAppointmentEndpoint))
	mux.Handle("POST /prescriptions", middleware.Auth(jwtSecret)(createPrescriptionEndpoint))
	mux.Handle("GET /prescriptions", middleware.Auth(jwtSecret)(getPatientPrescriptionEndpoint))
	mux.Handle("POST /payments", middleware.Auth(jwtSecret)(createPaymentEndpoint))

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: middleware.Logger(mux),
	}
	log.Println("Server running on :" + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}
