package day_04

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  4,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}
