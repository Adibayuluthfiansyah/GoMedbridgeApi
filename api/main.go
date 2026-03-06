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

	userRepo := repository.NewPostgresUserRepository(db)

	jwtSecret := cfg.JWTSecret

	userUsecase := usecase.NewUserUsecase(userRepo, jwtSecret)
	userHandler := handler.NewUserHandler(userUsecase)

	updateProfileEndpoint := http.HandlerFunc(userHandler.UpdateProfile)

	profileEndpoint := http.HandlerFunc(userHandler.GetProfile)

	//route
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.Handle("PUT /profile", middleware.Auth(jwtSecret)(updateProfileEndpoint))
	mux.Handle("GET /profile", middleware.Auth(jwtSecret)(profileEndpoint))
	mux.HandleFunc("GET /doctors", userHandler.GetDoctors)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: middleware.Logger(mux),
	}
	log.Println("Server running on :" + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}
