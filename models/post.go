package models

import (
	"database/sql"
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

// Aux post structure for posts
type AuxPost struct {
	ID     sql.NullInt64  // id of post
	IDUser sql.NullInt64  // id of user
	Text   sql.NullString // text of post
	Date   sql.NullTime   // date of post
}

// AddPost add post of the "posts" table
//  @param1 (post): structure variable "Post"
//
//  @return1 (err): error variable
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
//  @return1 (posts): post slice
//  @return2 (err): error variable
func GetPostsByUserID(id int) (posts []Post, err error) {
	var (
		auxPost AuxPost
		aux     Post
		date    sql.NullString
	)

	rows, err := DB.Query("SELECT date, text FROM posts WHERE user_id = ? ORDER BY date DESC", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&date, &auxPost.Text)
		if err != nil {
			return
		}

		aux.Date, err = time.Parse(time.RFC3339, "0000-01-01T00:00:00Z")
		if err != nil {
			return
		}

		if date.Valid {
			aux.Date, err = time.Parse(time.RFC3339, date.String)
			if err != nil {
				return
			}
		}

		aux.Text = "NULL"

		if auxPost.Text.Valid {
			aux.Text = auxPost.Text.String
		}

		posts = append(posts, aux)
	}

	return
}
