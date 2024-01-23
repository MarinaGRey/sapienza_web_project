package database

// remove all comments of a banned user
func (db *appdbimpl) RemoveComments(user User, banned User) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE UserID = ? AND PhotoID IN (SELECT PhotoID FROM photos WHERE UserID = ?)", banned.UserId, user.UserId)
	if err != nil {
		return err
	}

	return nil
}
