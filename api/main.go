package main

import (
	"log"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/config"
	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/internal/infrastructure/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.RunMigrations(db)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}
	log.Println("Server running on :" + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}
