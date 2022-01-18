package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Post structure for posts
type Post struct {
	ID        int       // id of post
	IDUser    int       // id of user
	Text      string    // text of post
	CreatedAt time.Time // post creation date
}

// AddPost add post of the "posts" table
//  @param1 (post): structure variable "Post"
//
//  @return1 (err): error variable
func AddPost(post Post) (err error) {
	_, err = DB.Exec("INSERT INTO posts (user_id, text) VALUES (?, ?)", post.IDUser, post.Text)
	if err != nil {
		return
	}

	return
}

// GetPostsByUserID get posts of user
//  @param1 (id): id of user
//
//  @return1 (posts): post slice
//  @return2 (err): error variable
func GetPostsByUserID(id int) (posts []Post, err error) {
	var (
		aux  Post
		date string
		text sql.NullString
	)

	rows, err := DB.Query("SELECT created_at, text FROM posts WHERE user_id = ? ORDER BY created_at DESC", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&date, &text)
		if err != nil {
			return
		}

		aux.CreatedAt, err = time.Parse(time.RFC3339, date)
		if err != nil {
			return

		}

		if text.Valid {
			aux.Text = text.String
		}

		posts = append(posts, aux)
	}

	return
}
