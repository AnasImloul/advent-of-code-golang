package year2023

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/2023/day_01"
)

func Run(day int, part string) any {
	switch day {
	case 1:
		return day_01.Solver.Solve(part)
	default:

		return nil
	}
}
