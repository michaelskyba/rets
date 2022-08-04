package main

import (
	"database/sql"
	"fmt"
)

func addCommand(db *sql.DB, args []string) {
	if len(args) != 5 {
		usageError()
	}

	table := args[1]
	filename := args[2]
	queryDate := args[3]

	fmt.Println(table, filename, queryDate)
}
