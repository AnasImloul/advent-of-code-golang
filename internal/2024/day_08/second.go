package day_08

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/types"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
	"strings"
)

func (d Day) secondPart() any {

	grid := strings.Split(d.ReadInput(), "\n")
	groups := d.getGroups(grid)

	checker := func(point, point1, point2 types.Point[int]) bool { return utils.AreAligned(point, point1, point2) }

	res := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if d.isAntiNode(types.Point[int]{X: x, Y: y}, groups, checker) {
				res += 1
			}
		}
	}

	return res
}
