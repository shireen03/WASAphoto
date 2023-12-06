package database

func (db *appdbimpl) SetLike(like Like) error {
	_, err := db.c.Exec("INSERT INTO like (likeID, userID, photoID) VALUES (?, ?, ?, ?)", like.LikeId, like.UserID, like.PhotoID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveLike(like Like) error {
	_, err := db.c.Exec("DELETE FROM like WHERE likeID=?", like.LikeId)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) GetLikeCount(pic Photo) (int, error) {
	var likeNumb int
	err := db.c.QueryRow("SELECT COUNT(*) FROM like WHERE photoID=?", pic.PhotoID).Scan(&likeNumb)
	if err != nil {
		return 0, err
	}
	return likeNumb, nil
}

func (db *appdbimpl) IsLiked(like Like) (bool, error) {
	var isliked bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM like WHERE userID=? AND photoID=?)", like.UserID, like.PhotoID).Scan(&isliked)
	if err != nil {
		return false, err
	}
	return isliked, nil

}
