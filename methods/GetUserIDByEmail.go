package forum

func GetUserIDByEmail(email string) (int, error) {
	var userID int
	// Query the database to retrieve the user ID associated with the provided email
	err := db.QueryRow("SELECT id FROM users WHERE email=?", email).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
