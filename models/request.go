package models

// <Request> structure for requests
//  @atr1 <Id_user_first  int>: id of first user
//  @atr2 <Id_user_second int>: id of second user
type Request struct {
	Id_user_first  int
	Id_user_second int
}

// <SendFriendRequest>     register friend requests in the "request" table
//  @param1 <req Request>: structure variable "Request"
//
//  @return1 <err error>:  error variable
func SendFriendRequest(req Request) (err error) {
	return
}

// <AnswerRequest> 		   add friend or delete the request
//  @param1 <req Request>: structure variable "Request"
//  @param2 <answ bool>:   answer of request true|false
//
//  @return1 <err error>:  error variable
func AnswerRequest(req Request, answ bool) (err error) {
	return
}

// <GetRequestsByUserName> get requests of user
//  @param1 <name string>: name of user
//
//  @return1 <p []Post>:   request slice
//  @return2 <err error>:  error variable
func GetRequestsByUserName(name string) (p []Request, err error) {
	return
}
