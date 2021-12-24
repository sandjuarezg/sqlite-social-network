package models

import (
	"database/sql"
	"errors"
)

// <Request> structure for requests
//  @atr1 <Id_user_first  int>: id of user who sent the request
//  @atr2 <Id_user_second int>: id of user who receives the request
type Request struct {
	Id_user_first  int
	Id_user_second int
}

// <SendFriendRequest>     register friend requests in the "request" table
//  @param1 <req Request>: structure variable "Request"
//
//  @return1 <err error>:  error variable
func SendFriendRequest(req Request) (err error) {
	if req.Id_user_first == req.Id_user_second {
		err = errors.New("that's your id user")
		return
	}

	row := DB.QueryRow("SELECT * FROM requests WHERE id_user_first = ? AND id_user_second = ?", req.Id_user_first, req.Id_user_second)
	if row.Scan().Error() != sql.ErrNoRows.Error() {
		err = errors.New("this request has already been sent")
		return
	}

	smt, err := DB.Prepare("INSERT INTO requests (id_user_first, id_user_second) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	_, err = smt.Exec(req.Id_user_first, req.Id_user_second)
	if err != nil {
		return
	}

	return
}

// <AnswerRequest> 		   add friend or delete the request
//  @param1 <req Request>: structure variable "Request"
//  @param2 <answ bool>:   answer of request true|false
//
//  @return1 <err error>:  error variable
func AnswerRequest(req Request, answ bool) (err error) {
	row, err := DB.Exec("DELETE from requests WHERE id_user_first = ? AND id_user_second = ?", req.Id_user_second, req.Id_user_first)
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

	if !answ {
		return
	}

	err = AddFriend(Friend{Id_user_first: req.Id_user_first, Id_user_second: req.Id_user_second})
	if err != nil {
		return
	}

	return
}

// <GetRequestsByIdUser>      get requests of user
//  @param1 <id int>:         id of user
//
//  @return1 <req []Request>: request slice
//  @return2 <err error>:     error variable
func GetRequestsByIdUser(id int) (req []Request, err error) {
	query := `
	SELECT 
		requests.id_user_first
		FROM 
			requests
		JOIN users ON users.id = requests.id_user_second
		WHERE 
			requests.id_user_second = ?
	`

	rows, err := DB.Query(query, id)
	if err != nil {
		return
	}
	defer rows.Close()

	var aux Request

	for rows.Next() {
		err = rows.Scan(&aux.Id_user_first)
		if err != nil {
			return
		}

		req = append(req, aux)
	}

	return
}
