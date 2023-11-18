package forum

import (
	"log"
	"net/http"
	"strconv"
)

func LikeDislikeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == http.MethodPost {
		// Retrieve the action (like/dislike) and postID from the form data
		action := r.PostFormValue("action")
		postID, err := strconv.Atoi(r.PostFormValue("postID"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		// Retrieve the user ID from the session
		userID, ok := GetAuthenticatedUserID(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Perform the appropriate action based on the 'action' value
		if action == "like" {
			err := LikePost(userID, postID) // Like the post
			if err != nil {
				log.Println("Error registering user:", err)
				http.Redirect(w, r, "/error/500", http.StatusSeeOther)
				return
			}
		} else if action == "dislike" {
			err := DislikePost(userID, postID) // Dislike the post
			if err != nil {
				log.Println("Error registering user:", err)
				http.Redirect(w, r, "/error/500", http.StatusSeeOther)
				return
			}
		} else {
			// If the action is neither 'like' nor 'dislike', return a bad request error
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		// Respond with a success message in JSON format
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": true}`))
	} else {
		// If the request method is not POST, return a method not allowed error
		http.Redirect(w, r, "/error/405", http.StatusSeeOther)
	}
}
