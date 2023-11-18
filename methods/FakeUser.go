package forum

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFakeUser() User {
	password := "testPassword"
	hashedPassword, err := GenerateTestPasswordHash(password)
	if err != nil {
		panic(err)
	}
	return User{
		ID:       12,
		Username: "test",
		Email:    "test@mysql.co.uk",
		Password: hashedPassword,
	}
}

func GenerateTestPasswordHash(password string) (string, error) {
	// Generate a bcrypt hash for the test password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
