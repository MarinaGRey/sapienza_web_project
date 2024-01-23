package database

// uncomment photo by deleting the comment from the databae
func (db *appdbimpl) UncommentPhoto(comment Comment) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE CommentID = ?", comment.CommentId)
	if err != nil {
		return err
	}

	return nil
}
