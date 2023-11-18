package forum

import (
	"net/http"
	"strconv"
)

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the authenticated user data
		userData := GetAuthenticatedUserData(r)

		// If the user is not authenticated, redirect to the login page
		if !userData.IsAuthenticated {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		postID, err := strconv.Atoi(r.FormValue("post_id"))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Retrieve the user ID from the session
		userID, ok := GetAuthenticatedUserID(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = DislikePost(userID, postID)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/error/405", http.StatusSeeOther)
}
