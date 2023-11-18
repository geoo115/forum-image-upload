package forum

func GetCommentsForPost(postID int) ([]Comment, error) {
	rows, err := db.Query(`SELECT comments.id, comments.post_id, comments.user_id, comments.content, comments.created_at, comments.updated_at, users.username
FROM comments  INNER JOIN users ON comments.user_id = users.id WHERE comments.post_id = ?`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		// var parentID sql.NullInt64
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.Username)
		if err != nil {
			return nil, err
		}

		// if parentID.Valid {
		// 	parentIDInt := int(parentID.Int64)
		// 	comment.ParentID = &parentIDInt
		// }

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for i, comment := range comments {
		// Get like count
		likeCount, err := GetLikeCountForComment(comment.ID)
		if err != nil {
			return nil, err
		}

		// Get dislike count
		dislikeCount, err := GetDislikeCountForComment(comment.ID)
		if err != nil {
			return nil, err
		}

		comments[i].LikeCount = likeCount
		comments[i].DislikeCount = dislikeCount
	}

	return comments, nil
}
