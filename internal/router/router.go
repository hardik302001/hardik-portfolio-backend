package router

import (
	"github.com/gorilla/mux"
	"github.com/hardik302001/hardik-portfolio-backend/internal/handler"
	"github.com/hardik302001/hardik-portfolio-backend/internal/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.CORS)
	r.Use(middleware.Logging)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/projects", handler.GetProjects).Methods("GET")
	api.HandleFunc("/about", handler.GetAbout).Methods("GET")
	api.HandleFunc("/contact", handler.PostContact).Methods("POST")
	api.HandleFunc("/health", handler.GetHealth).Methods("GET")
	api.HandleFunc("/profile", handler.GetProfile).Methods("GET")

	return r
}
