package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Read reads the content of the input file for the given year and day.
func Read(year, day int) string {
	// Format the day with leading zeros (e.g., "01" for day 1).
	filename := fmt.Sprintf("inputs/%d/day_%02d.txt", year, day)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read input file for year %d, day %02d: %v", year, day, err)
	}
	return strings.TrimSpace(string(data))
}

// ReadLines reads the input file for the given year and day and returns its content as a slice of lines.
func ReadLines(year, day int) []string {
	return strings.Split(Read(year, day), "\n")
}
