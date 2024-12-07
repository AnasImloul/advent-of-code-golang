package day_05

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
	"regexp"
	"strconv"
	"strings"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  5,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) readRules() [][]int {
	graph := make([][]int, 0)

	for _, line := range strings.Split(d.ReadInput(), "\n") {
		pattern := regexp.MustCompile(`(\d+)\|(\d+)`)

		matches := pattern.FindAllStringSubmatch(line, -1)

		if len(matches) == 0 {
			break
		}

		a, _ := strconv.Atoi(matches[0][1])
		b, _ := strconv.Atoi(matches[0][2])

		graph = append(graph, []int{a, b})
	}

	return graph
}

func (d Day) readSequences() [][]int {
	sequences := make([][]int, 0)
	for _, line := range strings.Split(d.ReadInput(), "\n") {
		if len(line) == 0 {
			continue
		}

		pattern := regexp.MustCompile(`(\d+)\|(\d+)`)

		matches := pattern.FindAllStringSubmatch(line, -1)

		if len(matches) != 0 {
			continue
		}

		sequence := utils.ToIntSlice(line, ",")

		sequences = append(sequences, sequence)
	}
	return sequences
}
