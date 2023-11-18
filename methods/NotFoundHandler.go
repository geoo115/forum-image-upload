package forum

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path
	if !(path == "/" || path == "/login" || path == "/register" || path == "/profile" || path == "/filter") {
		// http.Redirect(w, r, "/error/404", http.StatusNotFound)
		http.Redirect(w, r, "/error/404", http.StatusSeeOther)
		return true
	}
	return false
}
