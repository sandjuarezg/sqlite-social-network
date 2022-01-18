package models

import (
	"database/sql"
	"errors"
	"fmt"
)

// Request structure for requests
type Request struct {
	IDUserFirst  int // id of user who sent the request
	IDUserSecond int // id of user who receives the request
}

// Aux request structure for Requests
type AuxRequest struct {
	IDUserFirst  sql.NullInt64 // id of user who sent the request
	IDUserSecond sql.NullInt64 // id of user who receives the request
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

	err = DB.QueryRow(`
	SELECT 
		user_id_first 
		FROM 
			friends 
		WHERE 
			(user_id_first = ? AND user_id_second = ?) OR (user_id_second = ? AND user_id_first = ?)
	`, req.IDUserFirst, req.IDUserSecond, req.IDUserFirst, req.IDUserSecond).Scan()

	if err != sql.ErrNoRows {
		err = errors.New("they're already friends")
		return
	}

	err = DB.QueryRow("SELECT user_id_first FROM requests WHERE user_id_first = ? AND user_id_second = ?", req.IDUserFirst, req.IDUserSecond).Scan()
	if err != sql.ErrNoRows {
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

	var (
		auxRequest AuxRequest
		aux        Request
	)

	for rows.Next() {
		err = rows.Scan(&auxRequest.IDUserFirst)
		if err != nil {
			return
		}

		aux.IDUserFirst = -1

		if auxRequest.IDUserFirst.Valid {
			aux.IDUserFirst = int(auxRequest.IDUserFirst.Int64)
		}

		req = append(req, aux)
	}

	fmt.Println(req)

	return
}
