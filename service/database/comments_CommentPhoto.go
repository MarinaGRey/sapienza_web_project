package database

// allows commenting a photo by insering the necessary info in the database
func (db *appdbimpl) CommentPhoto(comment Comment) (Comment, error) {
	res, err := db.c.Exec(`INSERT INTO comments (PhotoID, UserID, Comment) VALUES (?, ?, ?)`, comment.PhotoId, comment.UserId, comment.Comment)

	if err != nil {
		return comment, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.CommentId = uint64(lastInsertId)

	return comment, nil
}
