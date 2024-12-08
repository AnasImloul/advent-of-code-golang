package day_08

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/types"
	"github.com/AnasImloul/advent-of-code-golang/internal/utils"
	"strings"
)

func (d Day) firstPart() any {

	grid := strings.Split(d.ReadInput(), "\n")
	groups := d.getGroups(grid)

	checker := func(point, point1, point2 types.Point[int]) bool {
		if !utils.AreAligned(point, point1, point2) {
			return false
		}

		d1, d2 := utils.DistanceSquared(point, point1), utils.DistanceSquared(point, point2)

		return d1 == 4*d2 || d2 == 4*d1
	}

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
