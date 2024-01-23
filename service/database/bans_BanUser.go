package database

// Database fuction that allows a user (banner) to ban another one (banned).
func (db *appdbimpl) BanUser(user User, banned User) error {

	_, err := db.c.Exec("INSERT INTO bans (UserID, BanUserID) VALUES (?, ?)", user.UserId, banned.UserId)
	if err != nil {
		return err
	}

	return nil
}
