package database

// uploads the photo inserting it in the database
func (db *appdbimpl) UploadPhoto(photo Photo) (Photo, error) {

	res, err := db.c.Exec(`INSERT INTO photos (UserId, File) VALUES (?, ?)`, photo.UserId, photo.File)

	if err != nil {
		return photo, err
	}

	lastInsertId, err := res.LastInsertId()

	if err != nil {
		return photo, err
	}

	photo.PhotoId = uint64(lastInsertId)

	return photo, nil
}
