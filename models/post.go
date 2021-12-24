package models

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// <Post> structure for posts
//  @atr1 <Id      int>: 	   id of post
//  @atr2 <Id_user int>: 	   id of user
//  @atr3 <Text    sttring>:   Text of post
//  @atr4 <Date    time.Time>: Date of post
type Post struct {
	Id      int
	Id_user int
	Text    string
	Date    time.Time
}

// <AddPost>              add post of the "posts" table
//  @param1 <p Post>: 	  structure variable "Post"
//
//  @return1 <err error>: error variable
func AddPost(p Post) (err error) {
	smt, err := DB.Prepare("INSERT INTO posts (id_user, text, date) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	_, err = smt.Exec(p.Id_user, p.Text, time.Now().Format(time.RFC3339))
	if err != nil {
		return
	}

	return
}

// <GetPostsByUserId>     get posts of user
//  @param1 <id int>:     id of user
//
//  @return1 <p []Post>:  post slice
//  @return2 <err error>: error variable
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
