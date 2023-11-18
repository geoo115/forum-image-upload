package forum

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(email, username, password string, db *sql.DB) error {
	// Generate a hashed password from the provided password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	// Insert the new user into the database
	query := "INSERT INTO users (email, username, password) VALUES (?, ?, ?)"

	// query := `INSERT INTO users \(email, username, password\) VALUES \(\?, \?, \?\)`
	_, err = db.Exec(query, email, username, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("error inserting new user: %v", err)
	}
	return nil
}
