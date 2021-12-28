package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// User structure for User
type User struct {
	ID       int    // id of user
	Email    string // email of user
	Username string // username of user
	Password string // password of user
}

// AddUser add user of the "users" table
//  @param1 (user): structure variable "User"

//  @return1 (err error): error variable
func AddUser(user User) (err error) {
	_, err = DB.Exec("INSERT INTO users (email, username, passwd) VALUES (?, ?, ?)", user.Email, user.Username, user.Password)
	if err != nil {
		err = errors.New("error: probably email already use")
		return
	}

	return
}

// LogIn user login
//  @param1 (email): email of user
//  @param2 (passwd): password of user
//
//  @return1 (user User): structure variable "User"
//  @return2 (err error): error variable
func LogIn(email, passwd string) (user User, err error) {
	row := DB.QueryRow("SELECT id, email, username, passwd FROM users WHERE email = ? AND passwd = ?", email, passwd)
	err = row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("User not found")
		}
		return
	}

	return
}

// DeleteAccount delete account of user
//  @rcvr1 (user): structure variable "User"
//
//  @return1 (err error): error variable
func (user User) DeleteAccount() (err error) {
	row, err := DB.Exec("DELETE from users WHERE id = ?", user.ID)
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

// GetSimilarUsersByUsername get similar usernames by username
//  @param1 (username): username
//
//  @return1 (users []User): user slice
//  @return2 (err error): error variable
func GetSimilarUsersByUsername(username string) (users []User, err error) {
	rows, err := DB.Query("SELECT id, username FROM users WHERE username LIKE ? ORDER BY id", fmt.Sprintf("%%%s%%", username))
	if err != nil {
		return
	}
	defer rows.Close()

	var aux User

	for rows.Next() {
		err = rows.Scan(&aux.ID, &aux.Username)
		if err != nil {
			return
		}

		users = append(users, aux)
	}

	return
}

// GetUsernameByUserID get username by id
//  @param1 (id): user id
//
//  @return1 (username string): username
//  @return2 (err error): error variable
func GetUsernameByUserID(id int) (username string, err error) {
	row := DB.QueryRow("SELECT username FROM users WHERE id = ?", id)
	err = row.Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("User not found")
		}
		return
	}

	return
}
