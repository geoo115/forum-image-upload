package forum

import (
	"database/sql"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	sessions     = make(map[string]int)
	userSessions = make(map[int]string)
	mu           sync.Mutex
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the email and password from the form values
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Prepare a query to retrieve the hashed password from the database
		stmt, err := db.Prepare("SELECT password FROM users WHERE email=?")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		defer stmt.Close()

		var storedPassword string
		// Query for the hashed password associated with the provided email
		err = stmt.QueryRow(email).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				// User not found
				tmpl, _ := template.ParseFiles("templates/login.html")
				tmpl.Execute(w, struct{ ErrorMessage string }{"Invalid email or password"})
				return
			} else {
				// Some other error occurred
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				log.Println(err)
				return
			}
		}

		// Compare the provided password with the stored hashed password
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			// Incorrect password
			tmpl, _ := template.ParseFiles("templates/login.html")
			tmpl.Execute(w, struct{ ErrorMessage string }{"Invalid email or password"})
			return
		}

		// Assuming the authentication was successful, get the user ID
		userID, err := GetUserIDByEmail(email)
		if err != nil {
			log.Println("Error getting user ID:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		log.Println("User ID:", userID)

		// Generate a new UUID for the session
		sessionID := uuid.New().String()

		mu.Lock()
		defer mu.Unlock()

		// Check if the user already has an active session
		if existingSessionID, ok := userSessions[userID]; ok {
			// If so, remove the existing session
			delete(userSessions, userID)
			log.Printf("Removed existing session for user %d\n", userID)
			// Also delete the session from the sessions map
			delete(sessions, existingSessionID)
		}

		// Store the session ID and user ID in their respective maps
		userSessions[userID] = sessionID
		sessions[sessionID] = userID

		// Store the session ID in a cookie with an expiration time
		expiration := time.Now().Add(24 * time.Hour) // 24 hours
		cookie := http.Cookie{
			Name:     "session-name",
			Value:    sessionID,
			Path:     "/",
			Expires:  expiration,
			HttpOnly: true,
			Secure:   true, // Enable only in production with HTTPS
		}

		http.SetCookie(w, &cookie)

		// Redirect the user to the home page after successful login
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If the request method is not POST, serve the login page
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl.Execute(w, nil)
}
