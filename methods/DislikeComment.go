package forum

import (
	"database/sql"
)

func DislikeComment(userID, commentID int) error {
	// Check if the user has already disliked the comment
	var currentLikeStatus bool
	err := db.QueryRow(`
		SELECT like_status FROM comments_likes_dislikes 
		WHERE user_id = ? AND comment_id = ?
	`, userID, commentID).Scan(&currentLikeStatus)

	switch {
	case err == sql.ErrNoRows:
		// User has not disliked the comment, insert the dislike
		_, err := db.Exec(`
			INSERT INTO comments_likes_dislikes (user_id, comment_id, like_status, count)
			VALUES (?, ?, ?, ?)
		`, userID, commentID, false, 1)
		return err
	case err != nil:
		return err
	default:
		if !currentLikeStatus {
			// User has already disliked the comment, remove the dislike
			_, err := db.Exec(`
				DELETE FROM comments_likes_dislikes
				WHERE user_id = ? AND comment_id = ?
			`, userID, commentID)
			return err
		}
		// User has liked the comment, update to dislike
		_, err := db.Exec(`
			UPDATE comments_likes_dislikes
			SET like_status = ?, count = count + 1
			WHERE user_id = ? AND comment_id = ?
		`, false, userID, commentID)
		return err
	}
}
