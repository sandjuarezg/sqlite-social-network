package models

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Post structure for posts
type Post struct {
	ID     int       // id of post
	IDUser int       // id of user
	Text   string    // text of post
	Date   time.Time // date of post
}

// AddPost add post of the "posts" table
//  @param1 (post): structure variable "Post"
//
//  @return1 (err error): error variable
func AddPost(post Post) (err error) {
	_, err = DB.Exec("INSERT INTO posts (user_id, text, date) VALUES (?, ?, ?)", post.IDUser, post.Text, time.Now().Format(time.RFC3339))
	if err != nil {
		return
	}

	return
}

// GetPostsByUserID get posts of user
//  @param1 (id): id of user
//
//  @return1 (posts []Post): post slice
//  @return2 (err error): error variable
func GetPostsByUserID(id int) (posts []Post, err error) {
	rows, err := DB.Query("SELECT date, text FROM posts WHERE user_id = ? ORDER BY date DESC", id)
	if err != nil {
		return
	}
	defer rows.Close()

	var (
		aux     Post
		content string
	)

	for rows.Next() {
		err = rows.Scan(&content, &aux.Text)
		if err != nil {
			return
		}

		aux.Date, err = time.Parse(time.RFC3339, string(content))
		if err != nil {
			return
		}

		posts = append(posts, aux)
	}

	return
}
