package forum

import (
	"database/sql"
)

func LikePost(userID, postID int) error {
	// Check if the user has already liked the post
	var currentLikeStatus bool
	err := db.QueryRow(`
		SELECT like_status FROM likes_dislikes 
		WHERE user_id = ? AND post_id = ?
	`, userID, postID).Scan(&currentLikeStatus)

	switch {
	case err == sql.ErrNoRows:
		// User has not liked the post, insert the like
		_, err := db.Exec(`
			INSERT INTO likes_dislikes (user_id, post_id, like_status, count)
			VALUES (?, ?, ?, ?)
		`, userID, postID, true, 1)
		return err
	case err != nil:
		return err
	default:
		if currentLikeStatus {
			// User has already liked the post, remove the like
			_, err := db.Exec(`
				DELETE FROM likes_dislikes
				WHERE user_id = ? AND post_id = ?
			`, userID, postID)
			return err
		}
		// User has disliked the post, update to like
		_, err := db.Exec(`
			UPDATE likes_dislikes
			SET like_status = ?, count = count + 1
			WHERE user_id = ? AND post_id = ?
		`, true, userID, postID)
		return err
	}
}
