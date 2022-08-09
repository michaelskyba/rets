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

// Add one, specific record to the given database
func pushData(db *sql.DB, table, queryDate string, record map[string]string) {
	fields := []string{}
	values := []string{}

	for field, value := range record {
		fields = append(fields, field)

		// Values have to be surrounded with quotes
		values = append(values, fmt.Sprintf("%q", value))
	}

	fieldString := strings.Join(fields, ",")
	valueString := strings.Join(values, ",")
	date := fmt.Sprintf("%q", queryDate)

	// We don't need to worry about SQL injection; only we will have access to
	// this script.
	statement := fmt.Sprintf("INSERT INTO %v (%v,rets_date) VALUES (%v,%v)", table, fieldString, valueString, date)
	_, err := db.Exec(statement)
	hdl(err)
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

			} else if columns.completed && columns.names[idx] == ignoreIndicator {
				// We're visiting an ignorable value during the record compilation
				continue
			}

			// We are looking at the field names declaration row (which should
			// be the third line in the file), so we need to keep special track of
			// them
			if !columns.completed {
				columns.names = append(columns.names, value)

			} else if value != "" {
				// If it's empty, it would show up as "", which doesn't fit into
				// DATE or INT. Instead, we'll just omit it so it will be left as its
				// empty valu in the SQL table.
				record[columns.names[idx]] = value
			}
		}

		// If we got this far, we must have just compiled the column names in
		// the values iteration
		if !columns.completed {
			columns.completed = true

		} else {
			// Add the compiled record to the database
			pushData(db, table, queryDate, record)
		}
	}
}
