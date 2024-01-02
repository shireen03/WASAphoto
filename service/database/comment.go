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

func (db *appdbimpl) GetComments(pic Photo) ([]Comment, error) {
	var comment []Comment
	rows, err := db.c.Query("SELECT userID, photoID, comment, photoUser FROM comment WHERE photoID = ?", pic.PhotoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var commenter Comment
		rows.Scan(&commenter.UserID, &commenter.PhotoID, &commenter.Text, &commenter.PhotoUserID)

		photoUsername, err := db.GetUsernameWithUserID(commenter.PhotoUserID)
		if err != nil {
			return nil, err
		}
		commenter.PhotoUsername = photoUsername

		Username, err := db.GetUsernameWithUserID(commenter.UserID)
		if err != nil {
			return nil, err
		}
		commenter.Username = Username

		comment = append(comment, commenter)
	}

	return comment, nil
}
