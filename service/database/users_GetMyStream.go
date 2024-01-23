package database

// Database function that gets the stream of a user (photos of people that are followed by the latter)
func (db *appdbimpl) GetStream(user User) ([]Photo, error) {
	var Photos []Photo

	rows, err := db.c.Query(`SELECT PhotoID FROM photos WHERE UserID IN (SELECT FollowUserID FROM followers WHERE UserID = ?) ORDER BY date DESC`, user.UserId)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset
	for rows.Next() {
		var photoId PhotoId
		var photo Photo

		err = rows.Scan(&photoId.PhotoId)
		if err != nil {
			return nil, err
		}

		photo, err = db.GetPhoto(photoId)
		if err != nil {
			return nil, err
		}

		Photos = append(Photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return Photos, nil
}
