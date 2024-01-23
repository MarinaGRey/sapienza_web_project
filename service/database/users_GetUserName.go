package database

import (
	"database/sql"
)

// retrieves the username in a type user structure
func (db *appdbimpl) GetUserName(user User) (User, error) {
	var userName User

	if err := db.c.QueryRow(`SELECT UserID, UserName FROM users WHERE UserID = ?`, user.UserId).Scan(&userName.UserId, &userName.UserName); err != nil {
		if err == sql.ErrNoRows {
			return userName, ErrUserDoesNotExist
		}
	}
	return userName, nil
}
