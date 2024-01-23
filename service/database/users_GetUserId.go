package database

import (
	"database/sql"
)

// retrieves the userid in a type user structure
func (db *appdbimpl) GetUserId(user User, requestUser User) (User, error) {

	if err := db.c.QueryRow(`SELECT UserID, UserName FROM users WHERE UserName = ? AND UserID NOT IN (SELECT UserID FROM Bans WHERE BanUserID = ?)`, user.UserName, requestUser.UserId).Scan(&user.UserId, &user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}

	return user, nil
}
