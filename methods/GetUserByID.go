package forum

import (
	"database/sql"
	"log"
)

func GetUserByID(userID int) (User, error) {
	var user User
	// Query user information by ID
	err := db.QueryRow("SELECT id, email, username FROM users WHERE id=?", userID).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with ID %d not found", userID)
			return user, err
		}
		log.Println("Error querying database:", err)
		return user, err
	}
	return user, nil
}
