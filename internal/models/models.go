package models

type Project struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	RepoURL     string   `json:"repo_url,omitempty"`
	LiveURL     string   `json:"live_url,omitempty"`
}

type AboutInfo struct {
	Name   string   `json:"name"`
	Bio    string   `json:"bio"`
	Skills []string `json:"skills"`
}

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
