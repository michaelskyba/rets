package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func addCommand(db *sql.DB, args []string) {
	if len(args) != 5 {
		usageError()
	}

	table := args[2]
	filename := args[3]
	queryDate := args[4]

	fmt.Println(table, filename, queryDate)

	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	hdl(err)

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "	")

		// Skip the first two rows and the very last row, which aren't normal
		// data entries
		if len(fields) == 1 {
			fmt.Println(fields)
		}
	}
}
