package day_06

func (d Day) secondPart() any {
	grid := d.readGrid()

	visited := d.markSafePositions(grid)

	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' || grid[i][j] == '^' || !visited[i][j] {
				continue
			}

			grid[i][j] = '#'
			_, loop := d.getSafePositions(grid)
			grid[i][j] = '.'

			if loop {
				res++
			}
		}
	}
	return res
}
