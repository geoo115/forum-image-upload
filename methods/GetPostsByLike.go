package forum

func GetPostsByLike(userID int, liked bool) ([]Post, error) {
	// Construct the query to get posts based on like status, including username
	query := `
		SELECT p.id, p.title, p.content, p.user_id, p.category, p.image_path,
			COALESCE(lc.like_count, 0) AS like_count, COALESCE(dc.dislike_count, 0) AS dislike_count,
			u.username
		FROM posts p
		LEFT JOIN (SELECT post_id, COUNT(*) AS like_count FROM likes_dislikes WHERE like_status = true GROUP BY post_id) lc ON p.id = lc.post_id
		LEFT JOIN (SELECT post_id, COUNT(*) AS dislike_count FROM likes_dislikes WHERE like_status = false GROUP BY post_id) dc ON p.id = dc.post_id
		JOIN users u ON p.user_id = u.id
		WHERE EXISTS (SELECT 1 FROM likes_dislikes ld WHERE ld.post_id = p.id AND ld.user_id = ? AND ld.like_status = ?)`

	rows, err := db.Query(query, userID, liked)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Category, &post.ImagePath, &post.LikeCount, &post.DislikeCount, &post.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
