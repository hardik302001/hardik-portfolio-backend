package handler

import (
	"net/http"

	"github.com/hardik302001/hardik-portfolio-backend/internal/models"
	"github.com/hardik302001/hardik-portfolio-backend/internal/response"
)

var about = models.AboutInfo{
	Name:   "Hardik Sharma",
	Bio:    "Software Developer",
	Skills: []string{"Go", "React", "TypeScript"},
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, about)
}
