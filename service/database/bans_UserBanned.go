package database

// [Util] Database fuction that checks if the requesting user was banned by another 'user'. Returns 'true' if is banned, 'false' otherwise
func (db *appdbimpl) UserBanned(user User, targetUser User) (bool, error) {
	var n int

	err := db.c.QueryRow("SELECT COUNT(*) FROM bans WHERE BanUserID = ? AND UserID = ?", user.UserId, targetUser.UserId).Scan(&n)

	if err != nil {
		return true, err
	}

	// If the counter is 1 then the user was banned
	if n == 1 {
		return true, nil
	}

	return false, nil
}
