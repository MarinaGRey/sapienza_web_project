package database

// unlikes the photo deleting the like
func (db *appdbimpl) UnlikePhoto(like Like) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE LikeID = ?", like.LikeId)
	if err != nil {
		return err
	}

	return nil
}
