package models

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// <CleanConsole>          clear console after 1 second
//
//  @return1 <err error>:  error variable
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

// <ScanRspnsWithMsgPrint>   print message and scan response
//  @param1 <msg string>: 	 message to display by console
//
//  @return1 <rspns string>: response of user
//  @return2 <err error>:    error variable
func ScanRspnsWithMsgPrint(msg string) (rspns string, err error) {
	fmt.Println(msg)
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return
	}
	rspns = string(aux)

	return
}

// <ReviewSqlMigration>    migration review
//
//  @return1 <err error>:  error variable
func ReviewSqlMigration() (err error) {
	_, err = os.Stat("./migration.sql")
	if os.IsNotExist(err) {
		err = errors.New("migration file not found")
		return
	}

	_, err = os.Stat("./social-network.db")
	if os.IsNotExist(err) {
		_, err = os.Create("./social-network.db")
		if err != nil {
			err = errors.New("error to create database")
			return
		}
	}

	content, err := os.ReadFile("./migration.sql")
	if err != nil {
		err = errors.New("error to read migration file")
		return
	}

	fdb, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		err = errors.New("error to open database")
		return
	}
	defer fdb.Close()

	_, err = fdb.Query("SELECT * from users")
	if err != nil {
		_, err = fdb.Exec(string(content))
		if err != nil {
			err = errors.New("error to execute migrations")
		}
	}

	return
}
