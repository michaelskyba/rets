package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
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

	tab := "	"

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	columns := strings.Split(scanner.Text(), tab)

	// Initial empty data array
	data := []int{}
	for i := 0; i < len(columns); i++ {
		data = append(data, 0)
	}

	// Fill in with data from the rest of the file
	rows := 0
	for scanner.Scan() {
		entryColumns := strings.Split(scanner.Text(), tab)
		rows++

		for idx, value := range entryColumns {
			if strings.TrimSpace(value) != "" {
				// The data value is not empty, so we will increment its count
				data[idx]++
			}
		}
	}

	// Compile output
	output := []string{}
	for idx, value := range data {
		percentage := math.Round(100 * float64(rows-value) / float64(rows))
		line := fmt.Sprintf("%v: %v/%v (%v%%) of entries were empty.",
		                    columns[idx], rows-value, rows, percentage)
		output = append(output, line)
	}

	err = os.WriteFile(os.Args[2], []byte(strings.Join(output, "\n")), 0644)
	hdl(err)
}
