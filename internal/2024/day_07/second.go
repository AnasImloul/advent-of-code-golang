package day_07

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
)

func (d Day) secondPart() any {
	var res int64 = 0

	// List of operations to use in secondPart (addition, multiplication, and concatenation)
	operations := []Operation{
		func(a, b int64) int64 { return a + b },                                             // addition
		func(a, b int64) int64 { return a * b },                                             // multiplication
		func(a, b int64) int64 { return a*int64(utils.Pow10(utils.NumberOfDigits(b))) + b }, // concatenation
	}

	for equation := range d.ReadEquations() {
		if checkPossible(equation, operations) { // Pass the list of operations
			res += int64(equation.Value)
		}
	}
	return res
}
