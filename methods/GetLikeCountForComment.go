package forum

func GetLikeCountForComment(commentID int) (int, error) {
	var likeCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM comments_likes_dislikes WHERE comment_id = ? AND like_status = 1`, commentID).Scan(&likeCount)
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}
