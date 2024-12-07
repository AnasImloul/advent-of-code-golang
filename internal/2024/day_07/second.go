package day_07

import (
	"fmt"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
)

func (d Day) secondPart() {
	var res int64 = 0

	// List of operations to use in secondPart (addition, multiplication, and concatenation)
	operations := []Operation{
		func(a, b int64) int64 { return a + b },             // addition
		func(a, b int64) int64 { return a * b },             // multiplication
		func(a, b int64) int64 { return concatenate(a, b) }, // concatenation
	}

	for equation := range d.ReadEquations() {
		if checkPossible(equation, operations) { // Pass the list of operations
			res += int64(equation.Value)
		}
	}
	fmt.Println(res)
}

func concatenate(a, b int64) int64 {
	return a*int64(utils.Pow10(utils.NumberOfDigits(b))) + b
}
