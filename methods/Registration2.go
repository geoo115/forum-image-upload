package forum

import (
	"log"
	"net/http"
	"text/template"
)

func Registration2(w http.ResponseWriter, r *http.Request) {
	// Create a User struct with default values
	User := User{
		SuccessfulRegistration: false,
	}

	// Check if the request method is POST or GET, return 400 Bad Request if not
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Redirect(w, r, "/error/400", http.StatusSeeOther)
		return
	}

	// Get the username, email, and password from the form values
	userN := r.FormValue("username")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	// Set RegistrationAttempted flag to true
	User.RegistrationAttempted = true

	// Check if the user already exists in the database
	exist, _ := UserExist(email, userN, db)

	// Parse the register.html template
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if exist {
		// If user already exists, set FailedRegister flag to true
		User.FailedRegister = true

		// Set a custom error message
		User.ErrorMessage = "User Name /Email is already taken"
		w.WriteHeader(http.StatusConflict) // Return a 409 Conflict status code
	} else {
		// If user doesn't exist, register the user and redirect to login page
		User.SuccessfulRegistration = true

		err := NewUser(email, userN, pass, db)
		if err != nil {
			log.Println("Error registering user:", err)
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			return
		}

		// Redirect to login page after successful registration
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err := tmpl.Execute(w, User)
	if err != nil {
		// If there's an error executing the template, return a 500 error and log the error
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println("Error executing template:", err)
		return
	}
}
