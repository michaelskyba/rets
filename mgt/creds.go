package main

import (
	"fmt"
	"os"
	"bufio"
)

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

	fmt.Println(username, password)
}
