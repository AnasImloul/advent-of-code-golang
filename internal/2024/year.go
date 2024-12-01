package year2024

import (
	"fmt"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_01"
)

func Run(day int, part string) {
	switch day {
	case 1:
		day_01.Solve(part)
	default:

		fmt.Printf("Day %d not implemented for 2024\n", day)
	}
}
