package day_01

import (
	"fmt"
	"strings"
)

// Map for literal number parsing
var numberLiterals = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

// parseLineWithLiterals checks each position in the line to find
// substrings that match numeric literals or digits and extracts
// the first and last numeric values.
func parseLineWithLiterals(line string) (int, int) {
	var left, right = -1, -1

	for i := 0; i < len(line); i++ {
		// Check if the current character is a digit
		if line[i] >= '0' && line[i] <= '9' {
			val := int(line[i] - '0')
			if left == -1 {
				left = val
			}
			right = val
			continue
		}

		// Check for a matching literal starting from position `i`
		for literal, val := range numberLiterals {
			if strings.HasPrefix(line[i:], literal) {
				if left == -1 {
					left = val
				}
				right = val
				break
			}
		}
	}

	return left, right
}

func (d Day01) secondPart() {
	var res = 0
	for _, line := range strings.Split(d.ReadInput(), "\n") {
		var left, right = parseLineWithLiterals(line)
		res += left*10 + right
	}
	fmt.Println(res)
}
