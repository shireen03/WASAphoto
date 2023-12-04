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
	"time"
)

type User struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
}
type Photo struct {
	PhotoID    uint64    `json:"photo_id"`
	UserID     uint64    `json:"user_id"`
	UploadTime time.Time `json:"uploadTime"`
	Date       string    `json:"date"`
	LikesNum   int       `json:"like_count"`
	CommentNum int       `json:"comment_count"`
	Comments   []string  `json:"comment_list"`
	Picture    []byte    `json:"pic"`
}

type Comment struct {
	CommentID uint64   `json:"comment_id"`
	UserID    uint64   `json:"user_id"`
	PhotoID   uint64   `json:"photo_id"`
	Text      string   `json:"text"`
	Comments  []string `json:"comments"`
}

type Like struct {
	LikeId  uint64 `json:"like_id"`
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
}

type Follow struct {
	UserID       uint64   `json:"user_id"`
	FollowID     uint64   `json:"follow_id"`
	FollowedID   uint64   `json:"followed_id"`
	FollowerList []uint64 `json:"follower_list"`
}

type Ban struct {
	UserID    uint64 `json:"user_id"`
	BanUserID uint64 `json:"ban_id"`
}

// this one is to get the array of pictures from users you follow
type Streamer struct {
	UserID         uint64   `json:"user_id"`
	StreamedPhotos []Stream `json:"streamedphotos"`
}

// this one describes the details of the stream which we will add into the array of pictures in Streamer
type Stream struct {
	Id             uint64 `json:"id"`
	UserID         uint64 `json:"user_id"`
	FollowedUserID uint64 `json:"followed_userID"`
	PhotoID        uint64 `json:"photo"`
	Date           string `json:"date"`
	LikeCount      int    `json:"comment_count"`
	CommentCount   int    `json:"like_count"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	LogIn(User) (User, error)
	SetUsername(string, User) (User, error)
	//GetUserId(uint64) (User, error)
	GetStream(User) ([]Stream, error)

	SetBan(Ban) error
	SetUnBan(Ban) error
	BanCheck(ppl1 uint64, ppl2 uint64) (bool, error)

	SetPic(Photo) error
	RemovePic(uint64) error

	SetLike(Like) error
	SetUnLike(Like) error
	GetLikeCount(pic Photo) (int, error)

	SetComment(Comment) error
	SetUnComment(Comment) error
	//GetCommentList(uint64)([]Comment,error)
	GetCommentCount(Photo) (int, error)
	RemoveBanComment(ban Ban) error

	SetFollow(Follow) error
	RemoveFollow(Follow) error
	GetFollowCount(Follow) (int, error)
	GetFollowingCount(Follow) (int, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='userTable';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		userDB := `CREATE TABLE user (
			UserID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT  
			);` //autoincrement cuz each userid is unique
		banDB := `CREATE TABLE ban (
			banID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userID INTEGER NOT NULL,
			bannedUserID INTEGER NOT NULL
			);`

		photoDB := `CREATE TABLE photos (
			photoID INTEGER NOT NULL PRIMARY KEY,
			userID INTEGER NOT NULL,
			date	TEXT NOT NULL,
			photo	BLOB
			FOREIGN KEY (userID) REFERENCES user(UserID),
			);`
		followDB := `CREATE TABLE follow (
				followID INTEGER NOT NULL PRIMARY KEY,
				userID INTEGER NOT NULL,
				toFollowID INTEGER NOT NULL,
				FOREIGN KEY (userID) REFERENCES user(UserID),
				);`

		likeDB := `CREATE TABLE like (
			likeID INTEGER NOT NULL PRIMARY KEY,
			userID INTEGER NOT NULL,
			photoID INTEGER NOT NULL,
			FOREIGN KEY (userID) REFERENCES user(UserID),
			FOREIGN KEY (photoID) REFERENCES photos(photoID)
			);`
		commentDB := `CREATE TABLE comment (
			commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			comment	TEXT NOT NULL,
			userID INTEGER NOT NULL,
			photoID INTEGER NOT NULL,
			FOREIGN KEY (userID) REFERENCES user(UserID),
			FOREIGN KEY (photoID) REFERENCES photos(photoID)
			);`

		_, err = db.Exec(userDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(commentDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(followDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(likeDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(banDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(photoDB)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
