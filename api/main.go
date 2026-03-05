package main

import (
	"log"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/delivery/http/handler"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/config"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/database"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/repository"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/usecase"
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

	database.RunMigrations(db)

	mux := http.NewServeMux()

	//handle regis
	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(repository.NewPostgresUserRepository(db)))
	mux.HandleFunc("/register", userHandler.Register)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}
	log.Println("Server running on :" + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}
