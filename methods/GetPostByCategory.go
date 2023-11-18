package forum

import "log"

func GetPostsByCategory(category string) ([]Post, error) {
	// Query posts with the specified category, including username
	rows, err := db.Query(`
        SELECT p.id, p.title, p.content, p.user_id, p.category, p.image_path, u.username 
        FROM posts p
        JOIN users u ON p.user_id = u.id
        WHERE category=?
    `, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	// Iterate over the results and populate the posts slice
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Category, &post.ImagePath, &post.Username)
		if err != nil {
			return nil, err
		}

		// Get like count for the post
		likeCount, err := GetLikeCountForPost(post.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		// Get dislike count for the post
		dislikeCount, err := GetDislikeCountForPost(post.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		post.LikeCount = likeCount
		post.DislikeCount = dislikeCount

		posts = append(posts, post)
	}

	// Check for any error during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
