package forum

func GetPostsByMyPost(userID int) ([]Post, error) {
	// Construct the query to get posts created by the user, including username
	query := `
		SELECT p.id, p.title, p.content, p.user_id, p.category, p.image_path,
			(SELECT COUNT(*) FROM likes_dislikes ld WHERE ld.post_id = p.id AND ld.like_status = true) AS like_count, 
			(SELECT COUNT(*) FROM likes_dislikes ld WHERE ld.post_id = p.id AND ld.like_status = false) AS dislike_count,
			u.username
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = ?`

	rows, err := db.Query(query, userID)
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
