package main

import (
	"database/sql"
	"fmt"
	forum "forum/methods"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err1 error
	db, err1 = sql.Open("sqlite3", "database.db")

	if err1 != nil {
		log.Fatal(err1.Error())
	}
	defer db.Close()

	// DEBUGGING - NO GO path
	// path := os.Getenv("GOPATH")
	// fmt.Println(path)

	// DEBUGGING
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Working Directory:", currentWorkingDir)

	forum.Init()
	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Add this line to serve the directory where you store uploaded images
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	http.HandleFunc("/", forum.HandleMain)
	http.HandleFunc("/register", forum.RegisterPage)
	http.HandleFunc("/register2", forum.Registration2)
	http.HandleFunc("/login", forum.LoginHandler)
	http.HandleFunc("/post-added", forum.AddPost)
	http.HandleFunc("/uploads", forum.UploadImage)
	http.HandleFunc("/logout", forum.LogoutHandler)
	http.HandleFunc("/filter", forum.FilterPosts)
	http.HandleFunc("/profile", forum.ProfilePage)
	http.HandleFunc("/add-comment", forum.AddCommentHandler)
	http.HandleFunc("/like", forum.LikePostHandler)
	http.HandleFunc("/dislike", forum.DislikePostHandler)
	http.HandleFunc("/like-dislike", forum.LikeDislikeHandler)
	http.HandleFunc("/like-comment", forum.LikeCommentHandler)
	http.HandleFunc("/dislike-comment", forum.DislikeCommentHandler)
	http.HandleFunc("/error/", forum.ErrorHandler)
	fmt.Println("Server started 🏁")
	fmt.Println("Listening at 👉 http://localhost:8000")
	fmt.Println("Ctrl+c to Close the Server ❌")

	// Consider using log.Fatal instead of log.Panic to handle server errors gracefully.
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
