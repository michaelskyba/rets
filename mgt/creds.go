package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// Handle errors
func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func getScriptText(username, password string) []byte {
	top := fmt.Sprintf(
`#!/bin/sh

# This file is supposed to kept off of GitHub so that Microsoft doesn't have
# access to the credentials.

[ -z "$1" ] &&
	echo "See the README.md for usage." &&
	exit 1

`)

	creds := fmt.Sprintf("export login='%v:%v'\n", username, password)
	bottom := fmt.Sprintln("./mgt/rets $@")

	return []byte(top + creds + bottom)
}

// Returns the filename (including path) of the pr script we're making
func getFilename() string {
	// e.g. "/home/michael/src/rets/mgt/creds"
	executable, err := os.Executable()
	hdl(err)

	// e.g. []string{"home", "michael", ..., "mgt", "creds"}
	path := strings.Split(executable, "/")

	// Get rid of the last two items to get the root directory
	root := path[0 : len(path)-2]

	return strings.Join(append(root, "pr"), "/")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("This will overwrite any previously-stored credentials.")
	fmt.Println("So, press Ctrl+c if you want to cancel.")

	fmt.Printf("Enter your RETS username.\n> ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Printf("Enter your RETS password.\n> ")
	scanner.Scan()
	password := scanner.Text()
	filename := getFilename()

	contents := getScriptText(username, password)
	err := os.WriteFile(filename, contents, 0755)
	hdl(err)
}
