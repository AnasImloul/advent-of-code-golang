package day_02

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/input"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  2,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func readReports() [][]int {
	var reports [][]int
	for line := range input.ReadLines(Solver.Year, Solver.Day) {
		if len(line) > 0 {
			reports = append(reports, utils.ToIntSlice(line, " "))
		}
	}
	return reports
}

func isReportSafe(report []int) bool {
	if !isMonotone(report) {
		return false
	}

	prev := report[0]

	for i := 1; i < len(report); i++ {
		if report[i] == prev || !isInRange(report[i], prev-3, prev+3) {
			return false
		}
		prev = report[i]
	}

	return true
}

func isInRange(num, a, b int) bool {
	return num >= a && num <= b
}

func isMonotone(report []int) bool {

	if len(report) < 2 {
		return true
	}

	var increasing = report[1] > report[0]

	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] && !increasing {
			return false
		} else if report[i] < report[i-1] && increasing {
			return false
		}
	}

	return true
}
