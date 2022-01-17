package models

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DB global database variable
var DB *sql.DB

func init() {
	err := ReviewSQLMigration()
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("sqlite3", "./social_network.db?_foreign_keys=ON")
	if err != nil {
		log.Fatal(err)
	}
}

// CleanConsole clear console after 1 second
//
//  @return1 (err): error variable
func CleanConsole() (err error) {
	time.Sleep(1 * time.Second)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return
	}

	return
}

// PrintMessageWithResponseScan print message and scan response
//  @param1 (msg): message to display by console
//
//  @return1 (rspns): response of user
//  @return2 (err): error variable
func PrintMessageWithResponseScan(msg string) (rspns string, err error) {
	fmt.Println(msg)
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return
	}
	rspns = string(aux)

	return
}

// ReviewSqlMigration migration review
//
//  @return1 (err):  error variable
func ReviewSQLMigration() (err error) {
	_, err = os.Stat("./migration.sql")
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New("migration file not found")
		}
		return
	}

	_, err = os.Stat("./social_network.db")
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create("./social_network.db")
			if err != nil {
				err = errors.New("error to create database")
				return
			}
		} else {
			return
		}
	}

	db, err := sql.Open("sqlite3", "./social_network.db")
	if err != nil {
		err = errors.New("error to open database")
		return
	}
	defer db.Close()

	_, err = db.Query("SELECT users.id, posts.id, friends.user_id_first, requests.user_id_first FROM users, posts, friends, requests")
	if err != nil {
		var content []byte

		content, err = os.ReadFile("./migration.sql")
		if err != nil {
			err = errors.New("error to read migration file")
			return
		}

		_, err = db.Exec(string(content))
		if err != nil {
			err = errors.New("error to execute migrations")
		}
	}

	return
}
