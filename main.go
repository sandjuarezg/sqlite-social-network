package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sandjuarezg/sqlite-social-network/models"
)

func main() {
	defer models.DB.Close()

	var (
		opc  int
		exit bool
	)

	for !exit {

		err := models.CleanConsole()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("0. Exit")
		fmt.Println("----------")
		fmt.Println("1. Log in")
		fmt.Println("2. Sing up")
		fmt.Scanln(&opc)

		err = models.CleanConsole()
		if err != nil {
			log.Println(err)
			continue
		}

		switch opc {
		case 0:

			exit = true
			fmt.Println(". . . .  B Y E  . . . .")

			err := models.CleanConsole()
			if err != nil {
				log.Println(err)
				continue
			}

		case 1:

			var back bool

			username, err := models.PrintMessageWithResponseScan("Enter username")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			password, err := models.PrintMessageWithResponseScan("Enter password")
			if err != nil {
				log.Println(err)
				continue
			}

			user, err := models.LogIn(username, password)
			if err != nil {
				log.Println(err)
				continue
			}

			for !back {

				err := models.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				opc = 0
				fmt.Printf("~ Welcome %s ~\n", user.Username)
				fmt.Println("0. Sign off")
				fmt.Println("1. Delete account")
				fmt.Println("-------------------")
				fmt.Println("2. Add post")
				fmt.Println("3. Add friend")
				fmt.Println("4. Delete friend")
				fmt.Println("-------------------")
				fmt.Println("5. Show your posts")
				fmt.Println("6. Show your friends")
				fmt.Println("7. Show your friend requests")
				fmt.Scanln(&opc)

				err = models.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				switch opc {
				case 0:

					back = true

					err := models.CleanConsole()
					if err != nil {
						log.Println(err)
						continue
					}

				case 1:
					var opc string

					fmt.Println("Are you sure you want to delete this account? y/n")
					fmt.Scanln(&opc)

					if opc != "y" {
						continue
					}

					back = true

					err = user.DeleteAccount()
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Account deleted successfully")

				case 2:

					text, err := models.PrintMessageWithResponseScan("Enter post text")
					if err != nil {
						log.Println(err)
						continue
					}

					err = models.AddPost(models.Post{IDUser: user.ID, Text: text})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post added successfully")

				case 3:

					username, err := models.PrintMessageWithResponseScan("Enter username to search")
					if err != nil {
						log.Println(err)
						continue
					}

					us, err := models.GetSimilarUsersByUsername(username)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					for _, v := range us {
						fmt.Printf("%d. %s\n", v.ID, v.Username)
					}

					var id int
					fmt.Println()
					fmt.Println("Enter user id to add")
					fmt.Scanln(&id)

					err = models.SendFriendRequest(models.Request{IDUserFirst: user.ID, IDUserSecond: id})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Request sent")

				case 4:

					id, err := models.PrintMessageWithResponseScan("Enter id user to delete")
					if err != nil {
						log.Println(err)
						continue
					}

					n, err := strconv.Atoi(id)
					if err != nil {
						log.Println(err)
						continue
					}

					err = models.DeleteFriend(models.Friend{IDUserFirst: user.ID, IDUserSecond: n})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Friend deleted")

				case 5:

					posts, err := models.GetPostsByUserID(user.ID)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(posts) == 0 {
						fmt.Println("you don't have posts yet")
						continue
					}

					fmt.Println("Your posts")
					fmt.Println()
					for _, v := range posts {
						fmt.Printf("%s - %s\n", v.Date, v.Text)
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 6:

					friends, err := models.GetFriendsByIDUser(user.ID)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(friends) == 0 {
						fmt.Println("you don't have friends yet")
						continue
					}

					for _, v := range friends {
						username, err := models.GetUsernameByUserID(v.IDUserFirst)
						if err != nil {
							log.Println(err)
							return
						}

						fmt.Printf("You and %s are friends since %s\n", username, v.Date)
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 7:

					requests, err := models.GetRequestsByIDUser(user.ID)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(requests) == 0 {
						fmt.Println("you don't have requests yet")
						continue
					}

					fmt.Println("Enter id of user")
					for _, v := range requests {
						username, err := models.GetUsernameByUserID(v.IDUserFirst)
						if err != nil {
							log.Println(err)
							return
						}

						fmt.Printf("%d. %s\n", v.IDUserFirst, username)
					}
					fmt.Scanln(&opc)

					var (
						ban  bool
						answ string
					)

					fmt.Println("Accept or reject friend request? a / r")
					fmt.Scanln(&answ)

					if answ == "a" {
						ban = true
					}

					err = models.AnswerRequest(models.Request{IDUserFirst: user.ID, IDUserSecond: opc}, ban)
					if err != nil {
						log.Println(err)
						continue
					}
				}

			}

		case 2:

			username, err := models.PrintMessageWithResponseScan("Enter username")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			password, err := models.PrintMessageWithResponseScan("Enter password")
			if err != nil {
				log.Println(err)
				continue
			}

			err = models.AddUser(models.User{Username: username, Password: password})
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("User added successfully")

		}

	}

}
