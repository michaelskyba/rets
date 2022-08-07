package main

import (
	"fmt"
	"os"
)

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func main() {
	if len(os.Args) != 3 {
		printError("Invalid usage. See gen_sql/README.md.")
	}
}
