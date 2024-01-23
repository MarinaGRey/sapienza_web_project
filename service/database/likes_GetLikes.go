package database

// list of users that liked a photo
func (db *appdbimpl) GetLikes(photo PhotoId) ([]Like, error) {
	var likes []Like
	var like Like

	rows, err := db.c.Query("SELECT * FROM likes WHERE PhotoID = ?", photo.PhotoId)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the likes in the resulset (likes of the photo with authors that didn't ban the requesting user).
	for rows.Next() {
		err = rows.Scan(&like.LikeId, &like.PhotoId, &like.UserId)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}
