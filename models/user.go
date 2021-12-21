package models

// <User> structure for User
//  @atr1 <Id       int>: 	 id of user
//  @atr2 <Username string>: username of user
//  @atr3 <Passwd   string>: password of user
type User struct {
	Id       int
	Username string
	Passwd   string
}

// <ExistUserByUsername>       user existence by name
//  @param1 <username string>: username of user
//
//  @return1 <ban bool>:       exist user true|false
//  @return2 <err error>:      error variable
func ExistUserByUsername(username string) (ban bool, err error) {
	return
}

// <AddUser>              add user of the "user" table
//  @param1 <u User>:     structure variable "User"
//
//  @return1 <err error>: error variable
func AddUser(u User) (err error) {
	return
}

// <LogIn>          		   user login
//  @param1 <username string>: username of user
//  @param2 <passwd string>:   password of user
//
//  @return1 <err error>:      error variable
func LogIn(username, passwd string) (u User, err error) {
	return
}

// <DeleteAccount>         delete account of user
//
//  @return1 <err error>:  error variable
func (u User) DeleteAccount() (err error) {
	return
}

// <GetUsernameByUserId>  	  	get username by user id
//  @param1 <id int>: 	  		user id
//
//  @return1 <username string>: username
//  @return2 <err error>:  	  	error variable
func GetUsernameByUserId(id int) (username string, err error) {
	return
}

// <GetUserIdByUsername>  	   get user id by username
//  @param1 <username string>: username
//
//  @return1 <id int>: 	  	   user id
//  @return2 <err error>:  	   error variable
func GetUserIdByUsername(username string) (id int, err error) {
	return
}
