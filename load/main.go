package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func usageError() {
	printError("Invalid usage. See the load/ README.md.")
}

func main() {
	config := mysql.Config{
		User:                 "rets",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "rets_db",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	hdl(err)

	err = db.Ping()
	hdl(err)

	if len(os.Args) < 2 {
		usageError()
	}

	switch os.Args[1] {
	case "add":
		addCommand(db, os.Args)

	default:
		usageError()
	}
}
