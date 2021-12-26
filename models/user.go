package models

import (
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// User: structure for User
type User struct {
	Id       int    // id of user
	Email    string // email of user
	Username string // username of user
	Passwd   string // password of user
}

// AddUser: add user of the "users" table
//  @param1 (u User): structure variable "User"

//  @return1 (err error): error variable
func AddUser(u User) (err error) {
	smt, err := DB.Prepare("INSERT INTO users (email, username, passwd) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	_, err = smt.Exec(u.Email, u.Username, u.Passwd)
	if err != nil {
		err = errors.New("error: probably email already use")
		return
	}

	return
}

// LogIn: user login
//  @param1 (email  string): email of user
//  @param2 (passwd string): password of user
//
//  @return1 (u User): structure variable "User"
//  @return2 (err error): error variable
func LogIn(email, passwd string) (u User, err error) {
	row := DB.QueryRow("SELECT id, email, username, passwd FROM users WHERE email = ? AND passwd = ?", email, passwd)
	err = row.Scan(&u.Id, &u.Email, &u.Username, &u.Passwd)
	if err != nil {
		err = errors.New("User not found")
		return
	}

	return
}

// DeleteAccount: delete account of user
//  @rcvr1 (u User): structure variable "User"
//
//  @return1 (err error): error variable
func (u User) DeleteAccount() (err error) {
	row, err := DB.Exec("DELETE from users WHERE id = ?", u.Id)
	if err != nil {
		return
	}

	n, err := row.RowsAffected()
	if err != nil {
		return
	}

	if n == 0 {
		err = errors.New("that user doesn't exist")
		return
	}

	return
}

// GetSimilarUsersByUsername: get similar usernames by username
//  @param1 (username string): username
//
//  @return1 (u []User): user slice
//  @return2 (err error): error variable
func GetSimilarUsersByUsername(username string) (u []User, err error) {
	rows, err := DB.Query("SELECT id, username FROM users WHERE username LIKE ? ORDER BY id", fmt.Sprintf("%%%s%%", username))
	if err != nil {
		return
	}
	defer rows.Close()

	var aux User

	for rows.Next() {
		err = rows.Scan(&aux.Id, &aux.Username)
		if err != nil {
			return
		}

		u = append(u, aux)
	}

	return
}

// GetUsernameByUserId: get username by id
//  @param1 (id int): user id
//
//  @return1 (username string): username
//  @return2 (err error): error variable
func GetUsernameByUserId(id int) (username string, err error) {
	row := DB.QueryRow("SELECT username FROM users WHERE id = ?", id)
	err = row.Scan(&username)
	if err != nil {
		err = errors.New("User not found")
		return
	}

	return
}
