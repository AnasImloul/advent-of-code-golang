package day_06

func (d Day) firstPart() any {
	grid := d.readGrid()

	res, _ := d.getSafePositions(grid)
	return res
}
