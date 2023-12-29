package database

import (
	"log"
	"strconv"

	"github.com/gofrs/uuid"
)

func (db *appdbimpl) LogUser(usr User) (User, error) {
	checkUser, err := db.CheckUserExist(usr)
	if err != nil {
		return User{}, err
	}
	if checkUser {
		user, err := db.GetUserWithUsername(usr.Username)
		if err != nil {
			return User{}, err
		}

		return user, nil
	} else {
		usr.UserID = usr.Username
		_, err := db.c.Exec("INSERT INTO user (username, userID) VALUES (?,?)", usr.Username, usr.UserID)
		if err != nil {
			return User{}, err
		}

		return usr, nil
	}

}
func (db *appdbimpl) LogtheUser(usr User) (string, error) {

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("couldnt generate userid : %v", err)
	}
	usr.UserID = id.String()

	_, err2 := db.c.Exec(`INSERT INTO user(username, userID) VALUES (?, ?)`,
		usr.Username, usr.UserID)
	if err2 != nil {
		return "ugh1", err
	}
	return usr.UserID, nil
}

func (db *appdbimpl) GetUserWithUsername(username string) (User, error) {
	var user User
	var tempUserID int64

	err := db.c.QueryRow("SELECT userID, username FROM user WHERE username=?", username).Scan(&tempUserID, &user.Username)
	if err != nil {
		return User{}, err
	}
	user.UserID = strconv.FormatInt(tempUserID, 10)

	return user, nil
}

func (db *appdbimpl) CheckUserExist(username User) (bool, error) {
	var userExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 from user WHERE username=?)", username.Username).Scan(&userExists)
	if err != nil {
		return false, err
	}
	return userExists, nil
}
func (db *appdbimpl) SetUsername(username string, user User) (User, error) {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE username=? AND userID=?", username, user.Username, user.UserID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// func (db *appdbimpl) GetProfile(usr User) (Profile, error) {
// 	var profile Profile
// 	var follow Follow
// 	profile.Username = usr.Username
// 	profile.UserID = usr.UserID

// 	follow.UserID = usr.UserID

// 	profile.FollowedCount, err = db.GetFollowCount(follow)

// }

// func (db *appdbimpl) GetStream(user User) ([]Stream, error) {
// 	//select photos from followed users with likes and comments count, in reverse chronological order
// 	rows, err := db.c.Query("SELECT p.photoID, p.userID, p.date "+
// 		"(SELECT COUNT(*) FROM like WHERE photoID = p.photoID) AS likeCount"+
// 		"(SELECT COUNT(*) FROM comment WHERE photoID = p.photoID) AS commentCount"+
// 		"FROM photos p INNER JOIN follow f ON p.userID = f.toFollowID WHERE"+
// 		"f.userID = ? ORDER BY p.date DESC", user.UserID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var streams []Stream
// 	for rows.Next() {
// 		var s Stream
// 		err := rows.Scan(&s.PhotoID, &s.FollowedUserID, &s.Date, &s.LikeCount, &s.CommentCount)
// 		if err != nil {
// 			return nil, err
// 		}
// 		s.UserID = user.UserID
// 		streams = append(streams, s)
// 	}

// 	return streams, nil
// }
