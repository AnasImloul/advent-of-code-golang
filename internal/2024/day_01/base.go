package day_01

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/input"
	"strconv"
	"strings"
)

type Day01 struct {
	day.Base
}

// Declare `d` as a global variable
var d = Day01{
	Base: day.Base{
		Year: 2024,
		Day:  1,
	},
}

func init() {
	// Initialize `d` once during package initialization
	d.FirstPart = d.firstPart
	d.SecondPart = d.secondPart
}

func Solve(part string) {
	d.Solve(part)
}

func readSequences() ([]int, []int) {

	var seq1 = []int{}
	var seq2 = []int{}

	for _, line := range input.ReadLines(d.Year, d.Day) {
		var nums = strings.Split(line, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		seq1 = append(seq1, num1)
		seq2 = append(seq2, num2)
	}
	return seq1, seq2
}
