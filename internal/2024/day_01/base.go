package day_01

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/input"
	"strconv"
	"strings"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  1,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func readSequences() ([]int, []int) {

	var seq1 []int
	var seq2 []int

	for line := range input.ReadLines(Solver.Year, Solver.Day) {
		var nums = strings.Split(line, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		seq1 = append(seq1, num1)
		seq2 = append(seq2, num2)
	}
	return seq1, seq2
}
