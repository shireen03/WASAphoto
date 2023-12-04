package database

func (db *appdbimpl) SetFollow(follow Follow) error {
	_, err := db.c.Exec("INSERT INTO follow (followID, userID, toFollowID) VALUES (?, ?, ?, ?)", follow.FollowID, follow.UserID, follow.FollowedID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveFollow(follow Follow) error {
	_, err := db.c.Exec("DELETE FROM follow WHERE followID=?", follow.FollowID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetFollowCount(follow Follow) (int, error) {
	var followCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM followers WHERE toFollowID = ? AND userID=?", follow.UserID, follow.FollowedID).Scan(&followCount)

	if err != nil {
		return 0, err
	}

	return followCount, nil
}

func (db *appdbimpl) GetFollowingCount(follow Follow) (int, error) {
	var followingCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM followers WHERE userID = ? AND toFollowID=?", follow.UserID, follow.FollowedID).Scan(&followingCount)

	if err != nil {
		return 0, err
	}

	return followingCount, nil
}
