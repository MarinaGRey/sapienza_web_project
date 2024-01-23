package database

import (
	"database/sql"
)

// creates the user inserting all the info about it
func (db *appdbimpl) CreateUser(user User) (User, error) {
	var newUser User

	res, err := db.c.Exec("INSERT INTO users(UserName) VALUES (?)", user.UserName)
	if err != nil {
		if err := db.c.QueryRow("SELECT UserID, UserName FROM users WHERE UserName = ?", user.UserName).Scan(&newUser.UserId, &newUser.UserName); err != nil {
			if err == sql.ErrNoRows {
				return user, ErrUserDoesNotExist
			}
		}
		return newUser, nil
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return user, err
	}

	newUser.UserId = uint64(lastInsertID)
	newUser.UserName = user.UserName

	return newUser, nil
}
