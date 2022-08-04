package main

import (
	"bufio"
	"os"
	"strings"
)

func hdl(err error, message string) {
	if err != nil {
		panic(message)
	}
}

func main() {
	if len(os.Args) != 3 {
		panic("Invalid usage. See the gen_fields README.md.")
	}

	inputFile, err := os.Open(os.Args[1])
	defer inputFile.Close()
	hdl(err, "Error: Failed to open input file.")

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	columns := strings.Split(scanner.Text(), "\t")

	scanner.Scan()
	values := strings.Split(scanner.Text(), "\t")

	output := []string{}

	for i, _ := range columns {
		column := columns[i]
		value := values[i]

		if value == "" {
			value = "gen.go: None"
		}

		output = append(output, "\n-- Column --------- :", column,
			"-- Value ---------- :", value)
	}

	data := []byte(strings.Join(output, "\n"))
	err = os.WriteFile(os.Args[2], data, 0644)
	hdl(err, "Error: Couldn't write to output file")
}
