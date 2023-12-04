package database

func (db *appdbimpl) SetBan(ban Ban) error {
	_, err := db.c.Exec("INSERT INTO comment (commentID, comment, userID, photoID) VALUES (?, ?, ?, ?)", c.CommentID, c.Text, c.UserID, c.PhotoID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetUnBan(unban Ban) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE commentID=?", c.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetBanList(ban Ban) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE userID=?", ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) BanCheck(userID1 uint64, userID2 uint64) (bool, error) {
	var isBan bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM bans WHERE bannerID=? AND bannedID=?)", userID1, userID2).Scan(&isBan)
	if err != nil {
		return false, err
	}
	return isBan, nil
}
