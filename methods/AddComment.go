package forum

import (
	"time"
)

func AddComment(postID, userID int, content string) error {
	// Insert a new comment into the database
	_, err := db.Exec("INSERT INTO comments (user_id,post_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		userID, postID, content, time.Now(), time.Now())
	return err
}
