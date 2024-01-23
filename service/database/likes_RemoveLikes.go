package database

// remove all likes of a ban user
func (db *appdbimpl) RemoveLikes(user User, banned User) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE UserID = ? AND PhotoID IN (SELECT PhotoID FROM photos WHERE UserID = ?)", banned.UserId, user.UserId)
	if err != nil {
		return err
	}

	return nil
}
