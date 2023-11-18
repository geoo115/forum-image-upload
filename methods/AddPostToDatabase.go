package forum

func AddPostToDatabase(category, title, content string, userID int, imagePath string) error {
	// Insert the post into the database
	_, err := db.Exec("INSERT INTO posts (title, content, user_id, category, image_path) VALUES (?, ?, ?, ?, ?)",
		title, content, userID, category, imagePath)
	return err
}
