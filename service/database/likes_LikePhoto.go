package database

// adds a like of a user to a photo
func (db *appdbimpl) LikePhoto(like Like) (Like, error) {

	res, err := db.c.Exec("INSERT INTO likes (PhotoID, UserID) VALUES (?, ?)", like.PhotoId, like.UserId)
	if err != nil {
		return like, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return like, err
	}

	like.LikeId = uint64(lastInsertId)

	return like, nil
}
