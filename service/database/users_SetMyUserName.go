package database

// changes the username with the same pk
func (db *appdbimpl) ChangeUsername(user User) error {

	_, err := db.c.Exec(`UPDATE users SET UserName = ? WHERE UserID = ?`, user.UserName, user.UserId)
	if err != nil {
		return err
	}

	return nil
}
