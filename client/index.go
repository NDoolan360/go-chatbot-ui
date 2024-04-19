package client

import (
	"html/template"
	"net/http"
)

// IndexTemplate represents the data structure for the index.tmpl.html template.
type IndexTemplate struct {
	Title       string
	Description string
}

// GetIndex handles the request for rendering the index page.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles(
		"client/templates/index.tmpl.html",
		"client/templates/head.tmpl.html",
	))

	err := indexTemplate.Execute(w, IndexTemplate{
		Title:       "Chatbot UI",
		Description: "A Description...",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
