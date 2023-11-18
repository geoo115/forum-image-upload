package forum

import "database/sql"

func DislikePost(userID, postID int) error {
	// Check if the user has already disliked the post
	var currentLikeStatus bool
	err := db.QueryRow(`
		SELECT like_status FROM likes_dislikes 
		WHERE user_id = ? AND post_id = ?
	`, userID, postID).Scan(&currentLikeStatus)
	switch {
	case err == sql.ErrNoRows:
		// User has not disliked the post, insert the dislike
		_, err := db.Exec(`
			INSERT INTO likes_dislikes (user_id, post_id, like_status, count)
			VALUES (?, ?, ?, ?)
		`, userID, postID, false, 1)
		return err
	case err != nil:
		return err
	default:
		if !currentLikeStatus {
			// User has already disliked the post, remove the dislike
			_, err := db.Exec(`
				DELETE FROM likes_dislikes
				WHERE user_id = ? AND post_id = ?
			`, userID, postID)
			return err
		}
		// User has liked the post, update to dislike
		_, err := db.Exec(`
			UPDATE likes_dislikes
			SET like_status = ?, count = count + 1
			WHERE user_id = ? AND post_id = ?
		`, false, userID, postID)
		return err
	}
}
