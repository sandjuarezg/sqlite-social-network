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
	Username string // username of user
	Password string // password of user
}

// Aux user structure for User
type AuxUser struct {
	ID       sql.NullInt64  // id of user
	Username sql.NullString // username of user
	Password sql.NullString // password of user
}

// AddUser add user of the "users" table
//  @param1 (user): structure variable "User"
//
//  @return1 (err): error variable
func AddUser(user User) (err error) {
	_, err = DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		err = errors.New("error: probably username already use")
		return
	}

	return
}

// LogIn user login
//  @param1 (username): username of user
//  @param2 (password): password of user
//
//  @return1 (user): structure variable "User"
//  @return2 (err): error variable
func LogIn(username, password string) (user User, err error) {
	var auxUser AuxUser

	row := DB.QueryRow("SELECT id, username, password FROM users WHERE username = ? AND password = ?", username, password)
	err = row.Scan(&auxUser.ID, &auxUser.Username, &auxUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("User not found")
		}
		return
	}

	user.ID = -1
	if auxUser.ID.Valid {
		user.ID = int(auxUser.ID.Int64)
	}

	user.Username = "NULL"
	if auxUser.Username.Valid {
		user.Username = auxUser.Username.String
	}

	user.Password = "NULL"
	if auxUser.Password.Valid {
		user.Password = auxUser.Password.String
	}

	return
}

// DeleteAccount delete account of user
//  @rcvr1 (user): structure variable "User"
//
//  @return1 (err): error variable
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
//  @return1 (users): user slice
//  @return2 (err): error variable
func GetSimilarUsersByUsername(username string) (users []User, err error) {
	rows, err := DB.Query("SELECT id, username FROM users WHERE username LIKE ? ORDER BY id", fmt.Sprintf("%%%s%%", username))
	if err != nil {
		return
	}
	defer rows.Close()

	var (
		aux     User
		auxUser AuxUser
	)

	for rows.Next() {
		err = rows.Scan(&auxUser.ID, &auxUser.Username)
		if err != nil {
			return
		}

		aux.ID = -1
		if auxUser.ID.Valid {
			aux.ID = int(auxUser.ID.Int64)
		}

		aux.Username = "NULL"
		if auxUser.Username.Valid {
			aux.Username = auxUser.Username.String
		}

		users = append(users, aux)
	}

	return
}

// GetUsernameByUserID get username by id
//  @param1 (id): user id
//
//  @return1 (username): username
//  @return2 (err): error variable
func GetUsernameByUserID(id int) (username string, err error) {
	var auxUser AuxUser

	err = DB.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&auxUser.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("User not found")
		}
		return
	}

	username = "NULL"

	if auxUser.Username.Valid {
		username = auxUser.Username.String
	}

	return
}
