package day_01

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
)

type Day01 struct {
	day.Base
}

func Solve(part string) {
	d := Day01{
		Base: day.Base{
			Year: 2023,
			Day:  1,
		},
	}
	d.FirstPart = d.firstPart
	d.SecondPart = d.secondPart
	d.Solve(part)
}
