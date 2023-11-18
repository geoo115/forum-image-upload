package forum

import "net/http"

func ClearSession(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   "session-name",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}
