package forum

import (
	"database/sql"
)

func LikeComment(userID, commentID int) error {
	// Check if the user has already liked the comment
	var currentLikeStatus bool
	err := db.QueryRow(`
		SELECT like_status FROM comments_likes_dislikes 
		WHERE user_id = ? AND comment_id = ?
	`, userID, commentID).Scan(&currentLikeStatus)

	switch {
	case err == sql.ErrNoRows:
		// User has not liked the comment, insert the like
		_, err := db.Exec(`
			INSERT INTO comments_likes_dislikes (user_id, comment_id, like_status, count)
			VALUES (?, ?, ?, ?)
		`, userID, commentID, true, 1)
		return err
	case err != nil:
		return err
	default:
		if currentLikeStatus {
			// User has already liked the comment, remove the like
			_, err := db.Exec(`
				DELETE FROM comments_likes_dislikes
				WHERE user_id = ? AND comment_id = ?
			`, userID, commentID)
			return err
		}
		// User has disliked the comment, update to like
		_, err := db.Exec(`
			UPDATE comments_likes_dislikes
			SET like_status = ?, count = count + 1
			WHERE user_id = ? AND comment_id = ?
		`, true, userID, commentID)
		return err
	}
}
