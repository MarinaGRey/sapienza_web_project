package database

// gets the user of a photo
func (db *appdbimpl) GetUserPhoto(photo PhotoId) (User, error) {
	var user User

	err := db.c.QueryRow("SELECT UserID FROM photos WHERE PhotoID = ?", photo.PhotoId).Scan(&user.UserId)

	return user, err
}
