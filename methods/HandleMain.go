package forum

import (
	"log"
	"net/http"
	"text/template"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if NotFoundHandler(w, r) {
		return
	}

	// Get the authenticated user data
	userData := GetAuthenticatedUserData(r)

	// Get posts from the database
	posts, err := GetPostsFromDatabase()
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		log.Fatal(err)
		return

	}

	// Iterate through posts to fetch comments for each post
	for i, post := range posts {
		comments, err := GetCommentsForPost(post.ID)
		if err != nil {
			// If there's an error fetching comments, return a 500 error and log the error
			http.Error(w, "Error fetching comments", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		// Assign the comments to the post
		posts[i].Comments = comments
	}

	// Create a PageData struct with authentication status, username, and posts
	data := PageData{
		IsAuthenticated: userData.IsAuthenticated,
		Username:        userData.Username,
		Posts:           posts,
	}

	// Parse the index.html template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// If there's an error parsing the template, return a 500 error and log the error
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		// If there's an error executing the template, return a 500 error and log the error
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
