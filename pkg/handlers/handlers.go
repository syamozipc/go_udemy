package handlers

import (
	"net/http"

	"github.com/syamozipc/go_udemy/config"
	"github.com/syamozipc/go_udemy/pkg/render"
)

// repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// Home is the handler for the home page
func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
