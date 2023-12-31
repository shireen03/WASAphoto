package database

func (db *appdbimpl) SetBan(ban Ban) error {
	_, err := db.c.Exec("INSERT INTO ban (username, bannedUsername) VALUES (?, ?)", ban.Username, ban.BanUsername)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetUnBan(unban Ban) error {
	_, err := db.c.Exec("DELETE FROM ban WHERE bannedUsername=? AND username=? ", unban.BanUsername, unban.Username)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) BanCheck(userID1 Ban) (bool, error) {
	var isBan bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM ban WHERE username=? AND bannedUsername=?)", userID1.Username, userID1.BanUsername).Scan(&isBan)
	if err != nil {
		return false, err
	}
	return isBan, nil
}
