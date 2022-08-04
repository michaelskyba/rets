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

var ignoreIndicator = "load/add.go:Ignore"

// Get a hash table of column names to ignore
func getIgnoreList(filename string) map[string]bool {
	file, err := os.Open(filename)
	hdl(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()

	ignore := map[string]bool{}
	for scanner.Scan() {
		ignore[scanner.Text()] = true
	}

	return ignore
}

func addCommand(db *sql.DB, args []string) {
	if len(args) != 6 {
		usageError()
	}

	table := args[2]
	inputFilename := args[3]
	ignoreFilename := args[4]
	ignore := getIgnoreList(ignoreFilename)
	queryDate := args[5]

	fmt.Println(table, inputFilename, ignoreFilename, ignore, queryDate)

	file, err := os.Open(inputFilename)
	hdl(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()

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

		for idx, value := range values {
			// We need null values so that values[idx] still matches
			// column.names[idx]
			if ignore[value] && !columns.completed {
				columns.names = append(columns.names, ignoreIndicator)
				continue

				// We're visiting an ignorable value during the record compilation
			} else if columns.completed && columns.names[idx] == ignoreIndicator {
				continue
			}

			// We are looking at the field names declaration row (which should
			// be the third line in the file), so we need to keep special track of
			// them
			if !columns.completed {
				columns.names = append(columns.names, value)

			} else {
				record[columns.names[idx]] = value
			}
		}

		// If we got this far, we must have just compiled the column names in
		// the values iteration
		if !columns.completed {
			columns.completed = true
		}

		fmt.Println(record)
	}

	fmt.Println(columns.names)
}
