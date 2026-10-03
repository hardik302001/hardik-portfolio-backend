package handler

import (
	"log"
	"net/http"

	"github.com/hardik302001/hardik-portfolio-backend/internal/models"
	"github.com/hardik302001/hardik-portfolio-backend/internal/response"
)

func PostContact(w http.ResponseWriter, r *http.Request) {
	var req models.ContactRequest
	if err := response.DecodeJSON(r, &req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Name == "" || req.Email == "" || req.Message == "" {
		response.WriteError(w, http.StatusBadRequest, "name, email, and message are required")
		return
	}

	log.Printf("Contact form: name=%s email=%s message=%s", req.Name, req.Email, req.Message)

	response.WriteJSON(w, http.StatusOK, map[string]string{"status": "message received"})
}
