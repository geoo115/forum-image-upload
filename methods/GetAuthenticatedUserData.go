package forum

import (
	"log"
	"net/http"
)

func GetAuthenticatedUserData(r *http.Request) struct {
	IsAuthenticated bool
	Username        string
} {
	// Retrieve the user ID from the session cookie
	userID, ok := GetAuthenticatedUserID(r)
	if !ok {
		return struct {
			IsAuthenticated bool
			Username        string
		}{false, ""}
	}

	user, err := GetUserByID(userID)
	if err != nil {
		log.Fatal(err)
	}

	return struct {
		IsAuthenticated bool
		Username        string
	}{true, user.Username}
}
