package database

func (db *appdbimpl) LogIn(usr User) (User,error) {
	_, err := db.c.Exec("INSERT INTO user (username) VALUES (?)", usr.Username)
	if err != nil {
		return err
	}
	return usr, nil
}

func (db *appdbimpl) setUsername(username string, user User) (User,error) {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE username=? AND userID=?", username, user.Username, user.UserID)
	if err != nil {
		return err
	}
	return user, nil
}


func (db *appdbimpl) GetStream(user User) ([]Stream, error) {
    //select photos from followed users with likes and comments count, in reverse chronological order
    rows, err := db.c.QueryRow("SELECT p.photoID, p.userID, p.date,
                    (SELECT COUNT(*) FROM like WHERE photoID = p.photoID) AS likeCount,          
                    (SELECT COUNT(*) FROM comment WHERE photoID = p.photoID) AS commentCount
              		FROM photos p
                	INNER JOIN follow f ON p.userID = f.toFollowID
              		WHERE f.userID = ?
             		ORDER BY p.date DESC", user.UserID)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var streams []Stream
    for rows.Next() {
        var s Stream
        // Scan the row into the Stream struct
        err := rows.Scan(&s.PhotoID, &s.FollowedUserID, &s.Date, &s.LikeCount, &s.CommentCount)
        if err != nil {
            return nil, fmt.Errorf("scanning stream row: %w", err)
        }
        s.UserID = user.UserID // Set the UserID to the requesting user's ID
        streams = append(streams, s)
    }

    // Check for any error during iteration
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("iterating stream rows: %w", err)
    }

    return streams, nil
}

