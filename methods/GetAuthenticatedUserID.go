package forum

import (
	"net/http"
)

func GetAuthenticatedUserID(r *http.Request) (int, bool) {
	cookie, err := r.Cookie("session-name")
	if err != nil {
		return 0, false
	}

	userID, ok := sessions[cookie.Value]
	return userID, ok
}
