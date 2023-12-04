package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetComment(c Comment) error {
	_, err := db.c.Exec("INSERT INTO comment (commentID, comment, userID, photoID) VALUES (?, ?, ?, ?)", c.CommentID, c.Text, c.UserID, c.PhotoID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetUnComment(c Comment) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE commentID=?", c.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBanComment(ban Ban) (Comment, error) {
	_, err := db.c.Exec("DELETE FROM comment WHERE userID=?", ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetCommentCount(pic Photo) (int, error) {
	var count int
	_, err := db.c.QueryRow("SELECT COUNT(*) FROM comment WHERE photoID=?", pic.PhotoID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return countNumb, nil
}

func (db *appdbimpl) GetCommentCounst(pic Photo) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM comment WHERE photoID=?", pic.PhotoID).Scan(&count)
	if err != nil {
		return 0, err // Return 0 and the error if it occurs
	}
	return count, nil // Return the count and nil if no errors occurred
}
