package handlers

import (
	"net/http"
	"xcasluw/golang-from-scratch/internal/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl")
}
