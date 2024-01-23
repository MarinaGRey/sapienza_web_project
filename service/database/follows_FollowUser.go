package database

// follows inserting the user
func (db *appdbimpl) FollowUser(user User, followed User) error {

	_, err := db.c.Exec("INSERT INTO followers (UserId, FollowUserId) VALUES (?, ?)", user.UserId, followed.UserId)
	if err != nil {
		return err
	}

	return nil
}
