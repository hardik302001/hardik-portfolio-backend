package handler

import (
	"net/http"

	"github.com/hardik302001/hardik-portfolio-backend/internal/models"
	"github.com/hardik302001/hardik-portfolio-backend/internal/response"
)

var projects = []models.Project{
	{
		ID:          "1",
		Title:       "Portfolio Website",
		Description: "My personal portfolio built with Go and React",
		TechStack:   []string{"Go", "React", "TypeScript"},
		RepoURL:     "https://github.com/hardik302001/hardik-portfolio-backend",
	},
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, projects)
}
