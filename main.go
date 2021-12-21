package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/sqlite-social-network/models"
)

func main() {
	var (
		opc  int
		exit bool
	)

	err := models.ReviewSqlMigration()
	if err != nil {
		log.Fatal(err)
	}

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

			username, err := models.ScanRspnsWithMsgPrint("Enter username")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			passwd, err := models.ScanRspnsWithMsgPrint("Enter password")
			if err != nil {
				log.Println(err)
				continue
			}

			u, err := models.LogIn(username, passwd)
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
				fmt.Printf("~ Welcome %s ~\n", u.Username)
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

					err = u.DeleteAccount()
					if err != nil {
						log.Println(err)
						continue
					}

					back = true

					fmt.Println()
					fmt.Println("Account deleted successfully")

				case 2:

					text, err := models.ScanRspnsWithMsgPrint("Enter post text")
					if err != nil {
						log.Println(err)
						continue
					}

					err = models.AddPost(models.Post{Id_user: u.Id, Text: text})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post added successfully")

				case 3:

					username, err = models.ScanRspnsWithMsgPrint("Enter username")
					if err != nil {
						log.Println(err)
						continue
					}

					id, err := models.GetUserIdByUsername(username)
					if err != nil {
						log.Println(err)
						continue
					}

					err = models.SendFriendRequest(models.Request{Id_user_first: u.Id, Id_user_second: id})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Request sent")

				case 4:

					username, err = models.ScanRspnsWithMsgPrint("Enter username")
					if err != nil {
						log.Println(err)
						continue
					}

					id, err := models.GetUserIdByUsername(username)
					if err != nil {
						log.Println(err)
						continue
					}

					err = models.DeleteFriend(models.Friend{Id_user_first: u.Id, Id_user_second: id})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Printf("%s now isn't your friend\n", username)

				case 5:

					posts, err := models.GetPostsByUserName(u.Username)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(posts) == 0 {
						fmt.Println("you don't have posts yet")
						continue
					}

					for _, v := range posts {
						fmt.Printf("%d. %s\n", v.Id, v.Text)
					}

					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 6:

					friends, err := models.GetFriendsByUsername(u.Username)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(friends) == 0 {
						fmt.Println("you don't have friends yet")
						continue
					}

					for _, v := range friends {
						username, err = models.GetUsernameByUserId(v.Id_user_second)
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

					requests, err := models.GetRequestsByUserName(u.Username)
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
						username, err = models.GetUsernameByUserId(v.Id_user_second)
						if err != nil {
							log.Println(err)
							return
						}

						fmt.Printf("%d. %s\n", v.Id_user_second, username)
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

					err = models.AnswerRequest(models.Request{Id_user_first: u.Id, Id_user_second: opc}, ban)
					if err != nil {
						log.Println(err)
						continue
					}
				}

			}

		case 2:

			username, err := models.ScanRspnsWithMsgPrint("Enter username")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			passwd, err := models.ScanRspnsWithMsgPrint("Enter password")
			if err != nil {
				log.Println(err)
				continue
			}

			err = models.AddUser(models.User{Username: username, Passwd: passwd})
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("User added successfully")

		}

	}

}
