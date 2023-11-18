package forum

import (
	"net/http"

	"github.com/google/uuid"
)

func CreateSession(w http.ResponseWriter, userID int) {
	sessionID := uuid.New().String()

	cookie := http.Cookie{
		Name:     "session-name",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   86400, // 24 hours
		HttpOnly: true,
		Secure:   true, // Enable only in production with HTTPS
	}

	http.SetCookie(w, &cookie)

	// Store the session ID and user ID in a map
	sessions[sessionID] = userID
}
