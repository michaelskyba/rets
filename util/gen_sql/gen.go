package main

import (
	"bufio"
	"fmt"
	"os"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func main() {
	if len(os.Args) != 3 {
		printError("Invalid usage. See gen_sql/README.md.")
	}

	file, err := os.Open(os.Args[1])
	hdl(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
