package database

// unfollows deleting the user
func (db *appdbimpl) UnfollowUser(user User, followed User) error {

	_, err := db.c.Exec("DELETE FROM followers WHERE UserID = ? AND FollowUserID = ?", user.UserId, followed.UserId)
	if err != nil {
		return err
	}

	return nil
}
