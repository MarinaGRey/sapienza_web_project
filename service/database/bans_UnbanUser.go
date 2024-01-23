package database

// Database fuction that allows a user (banner) to unban another one (banned).
func (db *appdbimpl) UnbanUser(user User, banned User) error {

	_, err := db.c.Exec("DELETE FROM bans WHERE UserID = ? AND BanUserID = ?", user.UserId, banned.UserId)
	if err != nil {
		return err
	}

	return nil
}
