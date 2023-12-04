package service

import (
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

type Banned struct {
	UserID  uint64 `json:"user_id"`
	BanList []Ban  `json:"banList"`
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
