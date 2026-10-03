package handler

import (
	"net/http"

	"github.com/hardik302001/hardik-portfolio-backend/internal/response"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
