package models

// <Post> structure for posts
//  @atr1 <Id      int>: 	  id of post
//  @atr2 <Id_user int>: 	  id of user
//  @atr3 <Text    sttring>:  Text of post
type Post struct {
	Id      int
	Id_user int
	Text    string
}

// <AddPost>              add post of the "post" table
//  @param1 <p Post>: 	  structure variable "Post"
//
//  @return1 <err error>: error variable
func AddPost(p Post) (err error) {
	return
}

// <GetPostsByUsername>    get posts of user
//  @param1 <name string>: name of user
//
//  @return1 <p []Post>:   post slice
//  @return2 <err error>:  error variable
func GetPostsByUserName(name string) (p []Post, err error) {
	return
}
