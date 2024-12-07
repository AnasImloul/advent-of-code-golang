package day

import "github.com/AnasImloul/advent-of-code-golang/internal/input"

type Solver interface {
	ReadInput() string
	Solve(part string) any
}

type PartFunc func() any

type Base struct {
	Year       int
	Day        int
	FirstPart  PartFunc
	SecondPart PartFunc
}

// Solve handles executing the correct part logic.
func (b Base) Solve(part string) any {
	switch part {
	case "first":
		if b.FirstPart != nil {
			return b.FirstPart()
		} else {
			panic("First part not implemented")
		}
	case "second":
		if b.SecondPart != nil {
			return b.SecondPart()
		} else {
			panic("Second part not implemented")
		}
	default:
		panic("Invalid part ('first', 'second')")
	}
}

// ReadInput reads the input for the given year and day.
func (b Base) ReadInput() string {
	return input.Read(b.Year, b.Day)
}

func (b Base) ReadLines() <-chan string {
	return input.ReadLines(b.Year, b.Day)
}
