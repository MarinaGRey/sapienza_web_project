package database

// Database function that retrieves the number of follows of a user
func (db *appdbimpl) GetFollowing(user User) ([]User, error) {
	var following []User

	rows, err := db.c.Query("SELECT FollowUserID FROM followers WHERE UserID = ?", user.UserId)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users followed by the requesting user)

	for rows.Next() {
		var followed User

		err = rows.Scan(&followed.UserId)
		if err != nil {
			return nil, err
		}
		following = append(following, followed)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}
