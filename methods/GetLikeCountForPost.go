package forum

func GetLikeCountForPost(postID int) (int, error) {
	var likeCount int
	err := db.QueryRow(`SELECT COUNT(*) FROM likes_dislikes WHERE post_id = ? AND like_status = 1`, postID).Scan(&likeCount)
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}
