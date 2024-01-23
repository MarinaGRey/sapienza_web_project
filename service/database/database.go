/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// errors
var ErrUserDoesNotExist = errors.New("Error User does not exist!")
var ErrPhotoDoesNotExists = errors.New("Error photo does not exist!")
var ErrCommentNotFound = errors.New("Error comment does not exist!")
var ErrLikeNotFound = errors.New("Error like does not exist!")
var ErrBanDoesNotExist = errors.New("Error ban does not exist!")
var ErrFollowDoesNotExist = errors.New("Error follow does not exist!")

// database interface
type AppDatabase interface {
	CreateUser(User) (User, error)
	GetUserName(User) (User, error)
	GetUserId(User, User) (User, error)
	ChangeUsername(User) error

	// GetUserProfile(username string) (User, error)

	GetStream(User) ([]Photo, error)

	BanUser(User, User) error
	UnbanUser(User, User) error
	UserBanned(User, User) (bool, error)

	UploadPhoto(Photo) (Photo, error)
	DeletePhoto(PhotoId) error
	CheckPhoto(PhotoId) (bool, error)
	GetPhotos(User) ([]Photo, error)
	GetPhoto(PhotoId) (Photo, error)
	GetUserPhoto(PhotoId) (User, error)
	GetPhotoFile(PhotoId) ([]byte, error)

	CommentPhoto(Comment) (Comment, error)
	UncommentPhoto(Comment) error
	GetComments(PhotoId) ([]Comment, error)
	RemoveComments(User, User) error

	FollowUser(User, User) error
	UnfollowUser(User, User) error
	GetFollowers(User) ([]User, error)
	GetFollowing(User) ([]User, error)

	LikePhoto(Like) (Like, error)
	RemoveLikes(User, User) error
	UnlikePhoto(Like) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (*appdbimpl, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE IF NOT EXISTS users (UserID INTEGER PRIMARY KEY AUTOINCREMENT,
			                                          UserName VARCHAR(16) NOT NULL UNIQUE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (PhotoID INTEGER PRIMARY KEY AUTOINCREMENT,
			                             UserID INTEGER NOT NULL,
			                             File BLOB NOT NULL,
			                             date DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
			                             FOREIGN KEY(UserID) REFERENCES users (UserID) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (LikeID INTEGER PRIMARY KEY AUTOINCREMENT,
		                                PhotoID INTEGER NOT NULL,
			                            UserID INTEGER NOT NULL,
			                            FOREIGN KEY(PhotoID) REFERENCES photos (PhotoID) ON DELETE CASCADE,
			                            FOREIGN KEY(UserID) REFERENCES users (UserID) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure likes: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
			                               PhotoID INTEGER NOT NULL,
			                               UserID INTEGER NOT NULL,
			                               Comment VARCHAR(100) NOT NULL,
			                               FOREIGN KEY(PhotoID) REFERENCES photos (photoiD) ON DELETE CASCADE,
			                               FOREIGN KEY(UserID) REFERENCES users (UserID) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure comments: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bans';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE bans (UserID INTEGER NOT NULL,
			                           BanUserID INTEGER NOT NULL,
			                           PRIMARY KEY (UserID, BanUserID),
			                           FOREIGN KEY(UserID) REFERENCES users (UserID) ON DELETE CASCADE,
			                           FOREIGN KEY(BanUserID) REFERENCES users (UserID) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure bans: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followers';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE followers (UserID INTEGER NOT NULL,
			                                FollowUserID INTEGER NOT NULL,
			                                PRIMARY KEY (UserID, FollowUserID),
			                                FOREIGN KEY(UserID) REFERENCES users (UserID) ON DELETE CASCADE,
			                                FOREIGN KEY(FollowUserID) REFERENCES users (UserID) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure followers: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
