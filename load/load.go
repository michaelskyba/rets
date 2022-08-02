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

type Record struct {
	id int
	retsClass string
	retsDate string
	AC string
}

func printRecords(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM records")
	hdl(err)
	defer rows.Close()

	for rows.Next() {
		r := Record{}
		rows.Scan(&r.id, &r.retsClass, &r.retsDate, &r.AC)

		fmt.Println(r)
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

	printRecords(db)
}
