package year2023

import (
	"advent-of-code-golang/internal/2023/day_01"
	"fmt"
)

func Run(day int, part string) {
	switch day {
	case 1:
		day_01.Solve(part)
	default:
		fmt.Printf("Day %d not implemented for 2023\n", day)
	}
}