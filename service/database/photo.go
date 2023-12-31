package database

// set pic
func (db *appdbimpl) SetPic(pic Photo) error {
	_, err := db.c.Exec("INSERT INTO photos ( userID, date, photo) VALUES (?, ?, ?)", pic.UserID, pic.Date, pic.Picture)
	if err != nil {
		return err // Return if error
	}
	return nil // void for no error
}

// remove pic
func (db *appdbimpl) RemovePic(picID Photo) error {
	// Delete from the photo table
	_, err := db.c.Exec("DELETE FROM photos WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err // Return error
	}

	// Delete from the like table
	_, err = db.c.Exec("DELETE FROM like WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err
	}

	// Delete from the comment table
	_, err = db.c.Exec("DELETE FROM comment WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err
	}

	return nil // void if no errors

}

func (db *appdbimpl) GetPhotoCount(userID string) (int, error) {
	var photoCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE userID = ?", userID).Scan(&photoCount)

	if err != nil {
		return 0, err
	}

	return photoCount, nil
}
