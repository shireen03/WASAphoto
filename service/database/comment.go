package database

func (db *appdbimpl) SetComment(c Comment) error {
	_, err := db.c.Exec("INSERT INTO comment (commentID, comment, userID, photoID) VALUES (?, ?, ?, ?)", c.CommentID, c.Text, c.UserID, c.PhotoID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveComment(c Comment) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE commentID=?", c.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBanComment(ban Ban) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE userID=?", ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetCommentCount(pic Photo) (int, error) {
	var countNumb int
	err := db.c.QueryRow("SELECT COUNT(*) FROM comment WHERE photoID=?", pic.PhotoID).Scan(&countNumb)
	if err != nil {
		return 0, err
	}
	return countNumb, nil
}
