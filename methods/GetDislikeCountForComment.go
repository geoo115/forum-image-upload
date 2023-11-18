package forum

func GetDislikeCountForComment(commentID int) (int, error) {
	var dislikeCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM comments_likes_dislikes WHERE comment_id = ? AND like_status = 0`, commentID).Scan(&dislikeCount)
	if err != nil {
		return 0, err
	}
	return dislikeCount, nil
}
