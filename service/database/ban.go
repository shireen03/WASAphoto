package database

func (db *appdbimpl) SetBan(ban Ban) error {
	_, err := db.c.Exec("INSERT INTO ban (userID, bannedUserID) VALUES (?, ?)", ban.UserID, ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetUnBan(unban Ban) error {
	_, err := db.c.Exec("DELETE FROM ban WHERE bannedUserID=? AND userID=? ", unban.BanUserID, unban.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) BanCheck(userID1 Ban) (bool, error) {
	var isBan bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM bans WHERE bannerID=? AND bannedUserID=?)", userID1.UserID, userID1.BanUserID).Scan(&isBan)
	if err != nil {
		return false, err
	}
	return isBan, nil
}
