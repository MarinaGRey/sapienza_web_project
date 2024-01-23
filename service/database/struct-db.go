package database

import "time"

type User struct {
	UserId   uint64 `json:"userid"`
	UserName string `json:"username"`
}

type Photo struct {
	PhotoId  uint64    `json:"photoid"`
	UserId   uint64    `json:"userid"`
	UserName string    `json:"username"`
	File     []byte    `json:"file"`
	Date     time.Time `json:"date"`
	Comments []Comment `json:"comments"`
	Likes    []Like    `json:"likes"`
}

type PhotoId struct {
	PhotoId uint64 `json:"photoid"`
	UserId  uint64 `json:"userid"`
}

type Comment struct {
	CommentId uint64 `json:"commentid"` // Identifier of a comment
	PhotoId   uint64 `json:"photoid"`   // Photo unique id
	UserId    uint64 `json:"userid"`    // User's unique id
	UserName  string `json:"username"`  // UserName of a user
	Comment   string `json:"comment"`   // Comment content
}

type Like struct {
	LikeId  uint64 `json:"likeid"`  // Comment id
	PhotoId uint64 `json:"photoid"` // Photo id
	UserId  uint64 `json:"userid"`  // User id
}
