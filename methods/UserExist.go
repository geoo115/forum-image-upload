package forum

import (
	"database/sql"
	"fmt"
)

func UserExist(email, username string, db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email=? OR username=?", email, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking if user exists: %v", err)
	}
	return count > 0, nil
}
