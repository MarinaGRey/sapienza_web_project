package api

import (
	"github.com/MarinaGRey/WASA/service/database"
	"time"
)

// these are the structures that will be used for this project
type User struct {
	UserId   uint64 `json:"userid"`
	UserName string `json:"username"`
}

type UserName struct {
	UserName string `json:"username"`
}

type Photo struct {
	PhotoId  uint64             `json:"photoid"`
	UserId   uint64             `json:"userid"`
	UserName string             `json:"username"`
	File     []byte             `json:"file"`
	Date     time.Time          `json:"date"`
	Comments []database.Comment `json:"comments"`
	Likes    []database.Like    `json:"likes"`
}

type PhotoId struct {
	PhotoId uint64 `json:"photoid"`
	UserId  uint64 `json:"userid"`
}

type Comment struct {
	CommentId uint64 `json:"commentid"`
	PhotoId   uint64 `json:"photoid"`
	UserId    uint64 `json:"userid"`
	UserName  string `json:"username"`
	Comment   string `json:"comment"`
}

type Like struct {
	LikeId  uint64 `json:"likeid"`
	PhotoId uint64 `json:"photoid"`
	UserId  uint64 `json:"userid"`
}

type CommentText struct {
	Comment string `json:"comment"`
}

type Profile struct {
	UserId    uint64           `json:"userid"`
	UserName  string           `json:"username"`
	Followers []database.User  `json:"followers"`
	Following []database.User  `json:"following"`
	Photos    []database.Photo `json:"photos"`
}

// their structure must be changed to a format that can be used in the database
func (u *User) ToDatabase() database.User {
	return database.User{
		UserId:   u.UserId,
		UserName: u.UserName,
	}
}

func (u *User) FromDatabase(user database.User) {
	u.UserId = user.UserId
	u.UserName = user.UserName
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		CommentId: c.CommentId,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		UserName:  c.UserName,
		Comment:   c.Comment,
	}
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.CommentId = comment.CommentId
	c.PhotoId = comment.PhotoId
	c.UserId = comment.UserId
	c.UserName = comment.UserName
	c.Comment = comment.Comment
}

func (l *Like) ToDatabase() database.Like {
	return database.Like{
		LikeId:  l.LikeId,
		PhotoId: l.PhotoId,
		UserId:  l.UserId,
	}
}

func (l *Like) FromDatabase(like database.Like) {
	l.LikeId = like.LikeId
	l.PhotoId = like.PhotoId
	l.UserId = like.UserId
}

func (p *PhotoId) ToDatabase() database.PhotoId {
	return database.PhotoId{
		PhotoId: p.PhotoId,
		UserId:  p.UserId,
	}
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		PhotoId:  p.PhotoId,
		UserId:   p.UserId,
		UserName: p.UserName,
		File:     p.File,
		Date:     p.Date,
		Comments: p.Comments,
		Likes:    p.Likes,
	}
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.PhotoId = photo.PhotoId
	p.UserId = photo.UserId
	p.UserName = photo.UserName
	p.File = photo.File
	p.Date = photo.Date
	p.Comments = photo.Comments
	p.Likes = photo.Likes
}

// The username of the user is valid if it is between 3 and 16 characters
func (u *UserName) isValid() bool {
	length := len([]rune(u.UserName))
	return 3 <= length && length <= 16
}
