package models

import (
	"database/sql"
	"errors"
)

// Request structure for requests
type Request struct {
	IDUserFirst  int // id of user who sent the request
	IDUserSecond int // id of user who receives the request
}

// SendFriendRequest register friend requests in the "request" table
//  @param1 (req): structure variable "Request"
//
//  @return1 (err): error variable
func SendFriendRequest(req Request) (err error) {
	if req.IDUserFirst == req.IDUserSecond {
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
	`, req.IDUserFirst, req.IDUserSecond, req.IDUserFirst, req.IDUserSecond)

	if row.Scan() != sql.ErrNoRows {
		err = errors.New("they're already friends")
		return
	}

	row = DB.QueryRow("SELECT user_id_first FROM requests WHERE user_id_first = ? AND user_id_second = ?", req.IDUserFirst, req.IDUserSecond)
	if row.Scan() != sql.ErrNoRows {
		err = errors.New("this request has already been sent")
		return
	}

	_, err = DB.Exec("INSERT INTO requests (user_id_first, user_id_second) VALUES (?, ?)", req.IDUserFirst, req.IDUserSecond)
	if err != nil {
		err = errors.New("error: probably user not found")
		return
	}

	return
}

// AnswerRequest add friend or delete the request
//  @param1 (req): structure variable "Request"
//  @param2 (answ): answer of request true|false
//
//  @return1 (err): error variable
func AnswerRequest(req Request, answ bool) (err error) {
	row, err := DB.Exec(`
	DELETE  
		FROM 
			requests 
		WHERE 
			(user_id_first = ? AND user_id_second = ?) OR (user_id_second = ? AND user_id_first = ?)
	`, req.IDUserSecond, req.IDUserFirst, req.IDUserSecond, req.IDUserFirst)
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

	if !answ {
		return
	}

	err = AddFriend(Friend{IDUserFirst: req.IDUserFirst, IDUserSecond: req.IDUserSecond})
	if err != nil {
		return
	}

	return
}

// GetRequestsByIDUser get requests of user
//  @param1 (id): id of user
//
//  @return1 (req): request slice
//  @return2 (err): error variable
func GetRequestsByIDUser(id int) (req []Request, err error) {
	rows, err := DB.Query("SELECT user_id_first FROM requests WHERE user_id_second = ?", id)
	if err != nil {
		return
	}
	defer rows.Close()

	var aux Request

	for rows.Next() {
		err = rows.Scan(&aux.IDUserFirst)
		if err != nil {
			return
		}

		req = append(req, aux)
	}

	return
}
