package forum

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/error/405", http.StatusSeeOther)
		return
	}

	// Parse the form data
	err := r.ParseMultipartForm(20 << 20) // 20 MB max size
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	// Retrieve the category, title, content, and image file from the form values
	category := r.FormValue("category")
	title := r.FormValue("title")
	content := r.FormValue("content")
	file, header, err := r.FormFile("image")

	fmt.Println(err)

	// Check if any of the fields are empty, excluding the image
	if category == " " || title == " " || content == " " {
		http.Redirect(w, r, "/error/400", http.StatusSeeOther)
		return
	}
	// Use a regular expression to check for non-empty content
	if ok, err := regexp.MatchString(`\S`, content); !ok || err != nil {
		http.Error(w, "🤔 CANNOT see your POST, make sure you write something cool!", http.StatusBadRequest)
		return
	}

	// Retrieve the user ID from the session cookie
	userID, isAuthenticated := GetAuthenticatedUserID(r)

	if !isAuthenticated {
		// If the user is not authenticated, redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var imagePath string

	// Check if an image was provided
	if file != nil {
		// Image provided, handle it
		ext := filepath.Ext(header.Filename)
		ext = strings.ToLower(ext)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".jfif" && ext != ".png" && ext != ".gif" && ext != ".webp" && ext != ".bmp" {
			http.Error(w, "🤔 Perhaps that file isn't an image that this forum supports?", http.StatusBadRequest)
			return
		}

		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		randFileName := make([]byte, 16)
		for i := range randFileName {
			randFileName[i] = charset[rng.Intn(len(charset))]
		}

		imagePath = "uploads/" + string(randFileName) + "_" + strconv.Itoa(userID) + ext
		outFile, err := os.Create(imagePath)
		fmt.Println("file of type " + ext + " saved to " + imagePath)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}
	}

	// Add the post to the database
	err = AddPostToDatabase(category, title, content, userID, imagePath)
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	// Redirect the user to the home page after successfully adding the post
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
