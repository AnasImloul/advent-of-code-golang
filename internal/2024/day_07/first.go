package day_07

import (
	"fmt"
)

func (d Day) firstPart() {
	var res int64 = 0

	// List of operations to use in firstPart (addition and multiplication only)
	operations := []Operation{
		func(a, b int64) int64 { return a + b }, // addition
		func(a, b int64) int64 { return a * b }, // multiplication
	}

	for equation := range d.ReadEquations() {
		if checkPossible(equation, operations) { // Pass the list of operations
			res += int64(equation.Value)
		}
	}
	fmt.Println(res)
}
