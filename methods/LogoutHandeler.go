package forum

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-name")
	if err == nil {
		delete(sessions, cookie.Value)
	}

	ClearSession(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
