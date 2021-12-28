package models

import (
	"database/sql"
	"errors"
	"time"
)

// Friend structure for friends
type Friend struct {
	IDUserFirst  int       // id of first user
	IDUserSecond int       // id of second user
	Date         time.Time // friendship start date
}

// AddFriend add friendship in the "friends" table
//  @param1 (frds): structure variable "Friend"
//
//  @return1 (err): error variable
func AddFriend(frds Friend) (err error) {
	if frds.IDUserFirst == frds.IDUserSecond {
		err = errors.New("that's your id user")
		return
	}

	row := DB.QueryRow(`
	SELECT 
		user_id_first 
		FROM 
			friends 
		WHERE 
			(user_id_first = ? AND user_id_second = ?) OR (user_id_second = ? AND user_id_first = ?)
	`, frds.IDUserFirst, frds.IDUserSecond, frds.IDUserFirst, frds.IDUserSecond)

	if row.Scan() != sql.ErrNoRows {
		err = errors.New("they're already friends")
		return
	}

	_, err = DB.Exec("INSERT INTO friends (user_id_first, user_id_second, date) VALUES (?, ?, ?)", frds.IDUserFirst, frds.IDUserSecond, time.Now().Format(time.RFC3339))
	if err != nil {
		return
	}

	return
}

// DeleteFriend delete friendship of the "friends" table
//  @param1 (frds): structure variable "Friend"
//
//  @return1 (err): error variable
func DeleteFriend(frds Friend) (err error) {
	if frds.IDUserFirst == frds.IDUserSecond {
		err = errors.New("that's your id user")
		return
	}

	row, err := DB.Exec(`
	DELETE  
		FROM 
			friends 
		WHERE 
			(user_id_first = ? AND user_id_second = ?) OR (user_id_second = ? AND user_id_first = ?)
	`, frds.IDUserFirst, frds.IDUserSecond, frds.IDUserFirst, frds.IDUserSecond)

	if err != nil {
		return
	}

	n, err := row.RowsAffected()
	if err != nil {
		return
	}

	if n == 0 {
		err = errors.New("that user isn't your friend")
		return
	}

	return
}

// GetFriendsByIDUser get friends of user
//  @param1 (id): id of user
//
//  @return1 (frds): friends slice
//  @return2 (err): error variable
func GetFriendsByIDUser(id int) (frds []Friend, err error) {
	query := `
	SELECT 
		friends.date, friends.user_id_first
		FROM 
			friends
		JOIN users AS uU ON uU.id = friends.user_id_first
		WHERE 
			friends.user_id_second = ?
	UNION ALL
	SELECT 
		friends.date, friends.user_id_second
		FROM 
			friends
		JOIN users AS uF ON uF.id = friends.user_id_second
		WHERE 
			friends.user_id_first = ?
	`

	rows, err := DB.Query(query, id, id)
	if err != nil {
		return
	}
	defer rows.Close()

	var (
		aux     Friend
		content string
	)

	for rows.Next() {
		err = rows.Scan(&content, &aux.IDUserFirst)
		if err != nil {
			return
		}

		aux.Date, err = time.Parse(time.RFC3339, string(content))
		if err != nil {
			return
		}

		frds = append(frds, aux)
	}

	return
}
