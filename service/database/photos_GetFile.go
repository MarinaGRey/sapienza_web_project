package database

// retrieves the file of the photo
func (db *appdbimpl) GetPhotoFile(photoId PhotoId) ([]byte, error) {
	var photo []byte

	err := db.c.QueryRow("SELECT File FROM photos WHERE PhotoID = ?", photoId.PhotoId).Scan(&photo)
	if err != nil {
		return photo, err
	}

	return photo, nil
}
