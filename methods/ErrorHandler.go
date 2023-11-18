package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	// Attempt to parse the error template

	//IMPORTANT use the following 1st line to run tests and the 2nd when runnnig the server
	// tmpl, err := template.ParseFiles("../templates/error.html")
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// If template parsing fails, redirect to a generic error page
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		return
	}

	var errorMessage string
	statusCode := http.StatusNotFound

	switch r.URL.Path {
	case "/error/404":
		errorMessage = "Page not found"
		statusCode = http.StatusNotFound
	case "/error/500":
		errorMessage = "Internal Server Error"
		statusCode = http.StatusInternalServerError
	case "/error/400":
		errorMessage = "Bad Request"
		statusCode = http.StatusBadRequest
	case "/error/405":
		errorMessage = "Method Not Allowed"
		statusCode = http.StatusMethodNotAllowed
	default:
		errorMessage = "Unknown Error"
	}

	data := struct {
		Error     string
		ErrorCode string
	}{
		Error:     errorMessage,
		ErrorCode: fmt.Sprintf("%d", statusCode),
	}

	// Set the response status code based on the error type
	w.WriteHeader(statusCode)

	// Execute the error template and write it to the response
	if err := tmpl.Execute(w, data); err != nil {
		// If template execution fails, respond with an internal server error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
