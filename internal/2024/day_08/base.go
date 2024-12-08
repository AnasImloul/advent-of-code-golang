package day_08

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/types"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  8,
	},
}

type isAntiNodeFunction func(point, point1, point2 types.Point[int]) bool

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) getGroups(grid []string) map[rune][]types.Point[int] {
	groups := make(map[rune][]types.Point[int])
	for y, line := range grid {
		for x, char := range line {
			if char == '.' {
				continue
			}
			groups[char] = append(groups[char], types.Point[int]{X: x, Y: y})
		}
	}
	return groups
}

func (d Day) isAntiNode(point types.Point[int], groups map[rune][]types.Point[int], function isAntiNodeFunction) bool {
	for _, group := range groups {
		for i := 0; i < len(group)-1; i++ {
			for j := i + 1; j < len(group); j++ {
				if function(point, group[i], group[j]) {
					return true
				}
			}
		}
	}
	return false
}
