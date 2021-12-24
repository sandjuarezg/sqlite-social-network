package models

import (
	"database/sql"
	"errors"
	"time"
)

// <Friend> structure for friends
//  @atr1 <Id_user_first  int>: 	  id of first user
//  @atr2 <Id_user_second int>: 	  id of second user
//  @atr3 <Date           time.Time>: friendship start date
type Friend struct {
	Id_user_first  int
	Id_user_second int
	Date           time.Time
}

// <AddFriend>     		   add friendship in the "friends" table
//  @param1 <frds Friend>: structure variable "Friend"
//
//  @return1 <err error>:  error variable
func AddFriend(frds Friend) (err error) {
	if frds.Id_user_first == frds.Id_user_second {
		err = errors.New("that's your id user")
		return
	}

	if frds.Id_user_first > frds.Id_user_second {
		aux := frds.Id_user_first
		frds.Id_user_first = frds.Id_user_second
		frds.Id_user_second = aux
	}

	row := DB.QueryRow("SELECT * FROM friends WHERE id_user_first = ? AND id_user_second = ?", frds.Id_user_first, frds.Id_user_second)
	if row.Scan().Error() != sql.ErrNoRows.Error() {
		err = errors.New("they're already friends")
		return
	}

	smt, err := DB.Prepare("INSERT INTO friends (id_user_first, id_user_second, date) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	_, err = smt.Exec(frds.Id_user_first, frds.Id_user_second, time.Now().Format(time.RFC3339))
	if err != nil {
		return
	}

	return
}

// <DeleteFriend>          delete friendship of the "friends" table
//  @param1 <frds Friend>: structure variable "Friend"
//
//  @return1 <err error>:  error variable
func DeleteFriend(frds Friend) (err error) {
	if frds.Id_user_first == frds.Id_user_second {
		err = errors.New("that's your id user")
		return
	}

	if frds.Id_user_first > frds.Id_user_second {
		aux := frds.Id_user_first
		frds.Id_user_first = frds.Id_user_second
		frds.Id_user_second = aux
	}

	row, err := DB.Exec("DELETE from friends WHERE id_user_first = ? and id_user_second = ?", frds.Id_user_first, frds.Id_user_second)
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

// <GetFriendsByIdUser>  	  get friends of user
//  @param1 <id int>: 	      id of user
//
//  @return1 <frds []Friend>: friends slice
//  @return2 <err error>:  	  error variable
func GetFriendsByIdUser(id int) (frds []Friend, err error) {
	query := `
	SELECT 
		friends.date, friends.id_user_first
		FROM 
			friends
		JOIN users AS uU ON uU.id = friends.id_user_first
		WHERE 
			friends.id_user_second = ?
	UNION ALL
	SELECT 
		friends.date, friends.id_user_second
		FROM 
			friends
		JOIN users AS uF ON uF.id = friends.id_user_second
		WHERE 
			friends.id_user_first = ?
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
		err = rows.Scan(&content, &aux.Id_user_first)
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
