package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	CreateTables()
}

func CreateTables() {
	// Create users table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			email TEXT UNIQUE,
			username TEXT,
			password TEXT
		)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create posts table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        user_id INTEGER,
        category TEXT,
				image_path TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id)
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create comments table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		user_id INTEGER,
		username TEXT,
		content TEXT,
		parent_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	//Create likes_dislikes table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes_dislikes (
		user_id INTEGER,
		post_id INTEGER,
		like_status BOOLEAN,
		count INTEGER DEFAULT 0, 
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(post_id) REFERENCES posts(id),
		PRIMARY KEY (user_id, post_id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create comments likes_dislikes table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments_likes_dislikes (
		user_id INTEGER,
		comment_id INTEGER,
		like_status BOOLEAN,
		count INTEGER DEFAULT 0,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(comment_id) REFERENCES comments(id),
		PRIMARY KEY (user_id, comment_id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create categories table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		)`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables created successfully")
}
