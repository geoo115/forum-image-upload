package forum

func GetDislikeCountForPost(postID int) (int, error) {
	var dislikeCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM likes_dislikes WHERE post_id = ? AND like_status = 0`, postID).Scan(&dislikeCount)
	if err != nil {
		return 0, err
	}
	return dislikeCount, nil
}
