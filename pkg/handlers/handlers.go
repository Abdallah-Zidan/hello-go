package handlers

import (
	"net/http"

	"github.com/Abdallah-Zidan/hello-go/pkg/config"
	"github.com/Abdallah-Zidan/hello-go/pkg/models"
	"github.com/Abdallah-Zidan/hello-go/pkg/render"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

// NewRepository create new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{app}
}

// NewHandlers sets repository for handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.app.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate("home.page.html", w, &models.TemplateData{
		StringMap: map[string]string{
			"appName":   "Hello Go",
			"pageTitle": "Home Page",
		},
	})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := repo.app.Session.GetString(r.Context(), "remote_ip")
	stringMap := map[string]string{
		"appName":   "Hello Go",
		"pageTitle": "About Page",
		"remoteIp":  remoteIp,
	}

	render.RenderTemplate("about.page.html", w, &models.TemplateData{StringMap: stringMap})
}
