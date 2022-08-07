package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Indices struct {
	SystemName    int
	MaximumLength int
	DataType      int
}

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message+" See gen_sql/README.md.")
	os.Exit(1)
}

func getIndices(input string) Indices {
	arr := [3]int{-1, -1, -1}
	columns := strings.Split(input, "	")

	for i, column := range columns {
		if column == "SystemName" {
			arr[0] = i

		} else if column == "MaximumLength" {
			arr[1] = i

		} else if column == "DataType" {
			arr[2] = i
		}
	}

	if arr[0] == -1 || arr[1] == -1 || arr[2] == -1 {
		printError("Invalid metadata input.")
	}

	return Indices{arr[0], arr[1], arr[2]}
}

func main() {
	if len(os.Args) != 3 {
		printError("Invalid usage.")
	}

	filename := os.Args[1]
	propertyName := os.Args[2]
	inputFound := false

	file, err := os.Open(filename)
	hdl(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	columnIndices := getIndices(scanner.Text())

	for scanner.Scan() {
		var maximumLength string
		var dataType string

		for i, value := range strings.Split(scanner.Text(), "	") {
			if i == columnIndices.SystemName {
				// We got to the row that the input property is on
				if value == propertyName {
					inputFound = true

				} else {
					break
				}

			} else if i == columnIndices.MaximumLength {
				maximumLength = value

			} else if i == columnIndices.DataType {
				dataType = value
			}
		}

		if inputFound {
			fmt.Println(maximumLength, dataType)
			os.Exit(0)
		}
	}

	if !inputFound {
		printError("Invalid property.")
	}
}
