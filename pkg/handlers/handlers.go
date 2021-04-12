package handlers

import (
	"net/http"

	"github.com/NganJason/hotel-booking/pkg/config"
	"github.com/NganJason/hotel-booking/pkg/models"
	"github.com/NganJason/hotel-booking/pkg/render"
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

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// HandleAbout handles get request for about page
func (repo *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again from handler"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

