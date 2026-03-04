package main

import (
	"log"
	"net/http"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSON(w, http.StatusOK, response.JSONResponse{
			Status:  "succes",
			Message: "service healthy",
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
