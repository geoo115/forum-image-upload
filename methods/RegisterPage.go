package forum

import (
	"log"
	"net/http"
	"text/template"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if NotFoundHandler(w, r) {
		return
	}
	// Parse the register.html template

	// IMPORTANT swap for tests
	// tmpl, err := template.ParseFiles("../templates/register.html")
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		// If there's an error parsing the template, return a 500 error and log the error
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Execute the template with no data
	tmpl.Execute(w, nil)
}
