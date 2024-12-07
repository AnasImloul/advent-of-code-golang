package day_01

import (
	"strings"
)

func (d Day) firstPart() any {
	var res = 0
	for _, line := range strings.Split(d.ReadInput(), "\n") {
		var left, right = parseLine(line)
		res += left*10 + right
	}
	return res
}

func parseLine(line string) (int, int) {
	var left, right = -1, -1
	// parse left to right
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			left = int(line[i] - '0')
			break
		}
	}

	// parse left to right
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			right = int(line[i] - '0')
			break
		}
	}

	return left, right
}
