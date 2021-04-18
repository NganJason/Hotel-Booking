package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/models"
	"github.com/NganJason/hotel-booking/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers assign a repository to Repo
func NewHandlers(r *Repository) {
	Repo = r
}

// HandleHome handles get request for home page
func (repo *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "index.page.html", &models.TemplateData{})
}

// HandleAbout handles get request for about page
func (repo *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

func (repo *Repository) HandleGenerals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (repo *Repository) HandleMajor(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "major.page.html", &models.TemplateData{})
}

func (repo *Repository) HandleSearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (repo *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted to search availability"))
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

func (repo *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}