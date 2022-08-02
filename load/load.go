package main

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config := mysql.Config{
		User: "rets",
		Passwd: "password",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "rets_db",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	hdl(err)

	err = db.Ping()
	hdl(err)

	fmt.Println("load.go: DB connection successful.")
}
