package models

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Post: structure for posts
type Post struct {
	Id     int       // id of post
	IDUser int       // id of user
	Text   string    // text of post
	Date   time.Time // date of post
}

// AddPost: add post of the "posts" table
//  @param1 (p Post): structure variable "Post"
//
//  @return1 (err error): error variable
func AddPost(p Post) (err error) {
	_, err = DB.Exec("INSERT INTO posts (id_user, text, date) VALUES (?, ?, ?)", p.IDUser, p.Text, time.Now().Format(time.RFC3339))
	if err != nil {
		return
	}

	return
}

// GetPostsByUserId: get posts of user
//  @param1 (id int): id of user
//
//  @return1 (p []Post): post slice
//  @return2 (err error): error variable
func GetPostsByUserId(id int) (p []Post, err error) {
	rows, err := DB.Query("SELECT date, text FROM posts WHERE id_user = ? ORDER BY date DESC", id)
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

		p = append(p, aux)
	}

	return
}
