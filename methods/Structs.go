package forum

import "time"

type Post struct {
	ID           int
	Title        string
	Content      string
	Username     string
	UserID       int
	Category     string
	ImagePath    string
	Comments     []Comment
	LikeCount    int
	DislikeCount int
	Timestamp    time.Time
}
type LikeDislike struct {
	ID        int
	UserID    int
	PostID    int
	Like      bool
	CreatedAt time.Time
}
type Comment struct {
	ID           int
	PostID       int
	UserID       int
	Username     string
	Content      string
	ParentID     *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikeCount    int
	DislikeCount int
}

type User struct {
	ID                     int
	Email                  string
	Username               string
	Password               string
	Attempted              bool
	RegistrationAttempted  bool
	FailedRegister         bool
	SuccessfulRegistration bool
	ErrorMessage           string
}

type PageData struct {
	IsAuthenticated bool
	Username        string
	Posts           []Post
}

type SessionData struct {
	UserID int
}

type ErrorData struct {
	Error     string
	ErrorCode string
}
