package day_03

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"regexp"
	"strconv"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  3,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) parseMul(mul string) int64 {
	pattern := regexp.MustCompile("\\d+")

	matches := pattern.FindAllString(mul, -1)

	a, _ := strconv.Atoi(matches[0])
	b, _ := strconv.Atoi(matches[1])
	return int64(a * b)
}
