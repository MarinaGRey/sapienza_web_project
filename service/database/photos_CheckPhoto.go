package database

// retieves the count of photo of a user and photo
func (db *appdbimpl) CheckPhoto(photo PhotoId) (bool, error) {
	var rows int

	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE (PhotoID = ? and UserID = ?)", photo.PhotoId, photo.UserId).Scan(&rows)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}
