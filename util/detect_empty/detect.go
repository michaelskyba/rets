package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 3 {
		panic("Invalid usage. See the detect_empty README.md.")
	}

	inputFile, err := os.Open(os.Args[1])
	defer inputFile.Close()
	hdl(err)

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	columns := strings.Split(scanner.Text(), "	")

	data := [][]string{}
	for i := 0; i < len(columns); i++ {
		data = append(data, []string{})
	}

	fmt.Println(data)
}
