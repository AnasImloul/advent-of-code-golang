package year2024

import (
	"fmt"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_01"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_02"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_03"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_04"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_05"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_06"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_07"
)

func Run(day int, part string) {
	switch day {
	case 1:
		day_01.Solver.Solve(part)
	case 2:
		day_02.Solver.Solve(part)
	case 3:
		day_03.Solver.Solve(part)
	case 4:
		day_04.Solver.Solve(part)
	case 5:
		day_05.Solver.Solve(part)
	case 6:
		day_06.Solver.Solve(part)
	case 7:
		day_07.Solver.Solve(part)
	default:
		fmt.Printf("Day %d not implemented for 2024\n", day)
	}
}
