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

	file, err := os.Open(os.Args[1])
	hdl(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	columnIndices := getIndices(scanner.Text())
}
