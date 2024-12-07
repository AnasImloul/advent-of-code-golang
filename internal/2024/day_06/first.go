package day_06

import (
	"fmt"
)

func (d Day) firstPart() {
	grid := d.readGrid()

	res, _ := d.getSafePositions(grid)
	fmt.Println(res)
}
