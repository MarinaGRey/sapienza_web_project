package database

// retrieves the info for an array of photos
func (db *appdbimpl) GetPhotos(user User) ([]Photo, error) {
	var photos []Photo
	var photoId PhotoId

	rows, err := db.c.Query("SELECT * FROM photos WHERE UserID = ? ORDER BY Date DESC", user.UserId)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the photos in the resultset
	for rows.Next() {
		var photo Photo

		err = rows.Scan(&photo.PhotoId, &photo.UserId, &photo.File, &photo.Date)
		if err != nil {
			return nil, err
		}

		photoId.PhotoId = photo.PhotoId

		userName, err := db.GetUserName(user)
		if err != nil {
			return nil, err
		}
		photo.UserName = userName.UserName

		comments, err := db.GetComments(photoId)
		if err != nil {
			return nil, err
		}
		photo.Comments = comments

		likes, err := db.GetLikes(photoId)
		if err != nil {
			return nil, err
		}
		photo.Likes = likes

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}
