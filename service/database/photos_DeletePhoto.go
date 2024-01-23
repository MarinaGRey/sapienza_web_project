package database

func (db *appdbimpl) DeletePhoto(photo PhotoId) error {

	// deletes the photo, likes and comments are cascade
	_, err := db.c.Exec(`DELETE FROM photos WHERE PhotoID = ?`, photo.PhotoId)
	if err != nil {
		return err
	}

	return nil
}
