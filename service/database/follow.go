package database

func (db *appdbimpl) SetFollow(follow Follow) error {
	_, err := db.c.Exec("INSERT INTO follow (followID, userID, toFollowID) VALUES (?, ?, ?)", follow.FollowID, follow.UserID, follow.FollowedID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveFollow(follow Follow) error {
	_, err := db.c.Exec("DELETE FROM follow WHERE toFollowID=?", follow.FollowedID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetFollowCount(follow Follow) (int, error) {
	var followCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE toFollowID = ? AND userID=?", follow.UserID, follow.FollowedID).Scan(&followCount)

	if err != nil {
		return 0, err
	}

	return followCount, nil
}

func (db *appdbimpl) GetFollowingCount(follow Follow) (int, error) {
	var followingCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE userID = ? AND toFollowID=?", follow.UserID, follow.FollowedID).Scan(&followingCount)

	if err != nil {
		return 0, err
	}

	return followingCount, nil
}

func (db *appdbimpl) IsFollowing(follow Follow) (bool, error) {
	var isFollowing bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM follow WHERE followID=? AND toFollowID=?)", follow.UserID, follow.FollowedID).Scan(&isFollowing)
	if err != nil {
		return false, err
	}
	return isFollowing, nil

}
