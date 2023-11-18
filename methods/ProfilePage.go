package forum

import (
	"log"
	"net/http"
	"text/template"
)

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	if NotFoundHandler(w, r) {
		return
	}
	// Retrieve the user ID from the session cookie
	userID, ok := GetAuthenticatedUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve user information from the database based on userID
	user, err := GetUserByID(userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
