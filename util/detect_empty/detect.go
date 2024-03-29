package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"math"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

type Entry struct {
	empty int
	name string
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
			if strings.TrimSpace(value) == "" {
				// The data value is empty, so we will increment its count
				data[idx]++
			}
		}
	}

	// Compile entries, which hold the name of the column and the number of
	// times it appeared empty. This makes it easy to sort later
	entries := []Entry{}
	for idx, value := range data {
		entries = append(entries, Entry{value, columns[idx]})
	}

	// Sort the entries based on how empty they are
	sort.Slice(entries, func (i, j int) bool {
		       return entries[i].empty > entries[j].empty
	})

	// Compile output
	output := []string{"Name" + tab + "# Empty" + tab + "% Empty"}
	for _, entry := range entries {
		percentage := math.Round(100 * float64(entry.empty) / float64(rows))
		line := fmt.Sprintf("%v%v%v%v%v",
		                    entry.name, tab, entry.empty, tab, percentage)
		output = append(output, line)
	}

	// End with a newline
	output = append(output, "")

	err = os.WriteFile(os.Args[2], []byte(strings.Join(output, "\n")), 0644)
	hdl(err)
}
