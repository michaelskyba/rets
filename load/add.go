package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

type Columns struct {
	names     []string
	completed bool
}

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

	columns := Columns{
		names:     []string{},
		completed: false,
	}

	// Iterate over xml file lines, adding each to the provided table
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "	")

		// Skip the first two rows and the very last row, which aren't normal
		// data entries
		if len(values) == 1 {
			continue
		}

		record := map[string]string{}
		last := len(values) - 1

		for idx, value := range values {
			// We don't care about the <DATA> and </DATA> columns
			if idx == 0 || idx == last {
				continue
			}

			// We are looking at the field names declaration row (which should
			// be the third line in the file), so we need to keep special track of
			// them
			if !columns.completed {
				columns.names = append(columns.names, value)

			} else {
				// We skipped the first column ("<DATA>") and are now one ahaed
				columnIndex := idx - 1
				record[columns.names[columnIndex]] = value
			}
		}

		// If we got this far, we must have just compiled the column names in
		// the values iteration
		if !columns.completed {
			columns.completed = true
		}

		fmt.Println(record)
	}
}
