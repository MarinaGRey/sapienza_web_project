package database

// retrieves all the info from a photo
func (db *appdbimpl) GetPhoto(photoId PhotoId) (Photo, error) {
	var photo Photo
	var user User

	err := db.c.QueryRow("SELECT * FROM photos WHERE PhotoID = ?", photoId.PhotoId).Scan(&photo.PhotoId, &photo.UserId, &photo.File, &photo.Date)
	if err != nil {
		return photo, err
	}

	user.UserId = photo.UserId

	userName, err := db.GetUserName(user)
	if err != nil {
		return photo, err
	}
	photo.UserName = userName.UserName

	comments, err := db.GetComments(photoId)
	if err != nil {
		return photo, err
	}
	photo.Comments = comments

	likes, err := db.GetLikes(photoId)
	if err != nil {
		return photo, err
	}
	photo.Likes = likes

	return photo, nil
}
