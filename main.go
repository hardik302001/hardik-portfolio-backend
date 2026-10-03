package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hardik302001/hardik-portfolio-backend/internal/router"
)

func main() {
	r := router.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
