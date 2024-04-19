package client

import (
	"html/template"
	"net/http"
)

// IndexTemplate represents the data structure for the index.html.tmpl template
type IndexTemplate struct {
	Title       string
	Description string
}

// GetIndex handles the request for rendering the index page
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles(
		"client/templates/index.html.tmpl",
		"client/templates/head.html.tmpl",
	))

	err := indexTemplate.Execute(w, IndexTemplate{
		Title:       "Chatbot UI",
		Description: "A Description...",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
