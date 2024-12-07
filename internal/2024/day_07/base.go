package day_07

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
	"strconv"
	"strings"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  7,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

type Equation struct {
	Value int
	Parts []int
}

func (d Day) ReadEquations() <-chan Equation {
	ch := make(chan Equation)

	go func() {
		defer close(ch)
		for line := range d.ReadLines() {
			if len(line) == 0 {
				continue
			}

			colonIndex := strings.IndexByte(line, ':')
			if colonIndex == -1 {
				continue
			}

			left := line[:colonIndex]
			right := line[colonIndex+2:]

			value, _ := strconv.Atoi(left)
			parts := utils.ToIntSlice(right, " ")

			ch <- Equation{
				Value: value,
				Parts: parts,
			}
		}
	}()

	return ch
}

// Operation Define the type for the operation function
type Operation func(int64, int64) int64

// Shared helper function to avoid duplication, now accepting a list of operations
func checkPossible(equation Equation, operations []Operation) bool {
	var helper func(equation Equation, current int64, index int) bool

	helper = func(equation Equation, current int64, index int) bool {
		if index == len(equation.Parts) {
			return int64(equation.Value) == current
		}

		if current > int64(equation.Value) {
			return false
		}

		// Iterate over the operations list and apply them during recursion
		for _, op := range operations {
			if helper(equation, op(current, int64(equation.Parts[index])), index+1) {
				return true
			}
		}
		return false
	}

	// Start recursion with the first part
	return helper(equation, int64(equation.Parts[0]), 1)
}
