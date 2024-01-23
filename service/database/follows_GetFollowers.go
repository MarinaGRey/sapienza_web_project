package database

func (db *appdbimpl) GetFollowers(user User) ([]User, error) {
	var followers []User

	rows, err := db.c.Query("SELECT UserID FROM followers WHERE FollowUserID = ?", user.UserId)
	if err != nil {
		return nil, err
	}
	
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that follow the requesting user)
	for rows.Next() {
		var follower User

		err = rows.Scan(&follower.UserId)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}
