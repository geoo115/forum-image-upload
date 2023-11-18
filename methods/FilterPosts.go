package forum

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func FilterPosts(w http.ResponseWriter, r *http.Request) {
	// Check for authentication and handle 404 if necessary
	if NotFoundHandler(w, r) {
		return
	}

	// Retrieve the category, createdPosts, and likedPosts from the form values
	category := r.FormValue("category")
	createdPosts := r.FormValue("createdPosts")
	likedPosts := r.FormValue("likedPosts")

	// Convert string values to boolean
	createdPostsBool, _ := strconv.ParseBool(createdPosts)
	likedPostsBool, _ := strconv.ParseBool(likedPosts)

	fmt.Printf("category: %s, createdPosts: %v, likedPosts: %v\n", category, createdPostsBool, likedPostsBool)

	// Retrieve the user ID from the session
	userID, ok := GetAuthenticatedUserID(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Retrieve the username from the database based on the userID
	username, err := GetUserByID(userID)
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	var posts []Post

	// Filter the posts based on the selected criteria
	switch {
	case category != "":
		posts, err = GetPostsByCategory(category)
	case createdPostsBool:
		posts, err = GetPostsByMyPost(userID)
	case likedPostsBool:
		posts, err = GetPostsByLike(userID, true)
	default:
		posts, err = GetPostsFromDatabase()
	}

	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	for i := range posts {
		comments, err := GetCommentsForPost(posts[i].ID)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}
		posts[i].Comments = comments
	}

	data := PageData{
		IsAuthenticated: true,
		Username:        username.Username,
		Posts:           posts,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}
}
