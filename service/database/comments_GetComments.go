package database

func (db *appdbimpl) GetComments(photo PhotoId) ([]Comment, error) {
	var comments []Comment
	var comment Comment
	var user User

	rows, err := db.c.Query("SELECT * FROM comments WHERE PhotoID = ?", photo.PhotoId)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the comments (comments of the photo with authors that didn't ban the requesting user).
	for rows.Next() {
		err = rows.Scan(&comment.CommentId, &comment.PhotoId, &comment.UserId, &comment.Comment)
		if err != nil {
			return nil, err
		}

		// obtain name of the user
		user.UserId = comment.UserId

		user, err := db.GetUserName(user)
		if err != nil {
			return nil, err
		}

		comment.UserName = user.UserName

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
