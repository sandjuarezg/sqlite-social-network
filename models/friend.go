package models

import (
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

// <AddFriend>     		   add friendship in the "friend" table
//  @param1 <frds Friend>: structure variable "Friend"
//
//  @return1 <err error>:  error variable
func AddFriend(frds Friend) (err error) {
	return
}

// <DeleteFriend>          delete friendship of the "friend" table
//  @param1 <frds Friend>: structure variable "Friend"
//
//  @return1 <err error>:  error variable
func DeleteFriend(frds Friend) (err error) {
	return
}

// <GetFriendsByUsername>  	  get friends of user
//  @param1 <name string>: 	  name of user
//
//  @return1 <frds []Friend>: friends slice
//  @return2 <err error>:  	  error variable
func GetFriendsByUsername(name string) (frds []Friend, err error) {
	return
}
