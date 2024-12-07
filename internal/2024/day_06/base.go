package day_06

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"strings"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  6,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) getPatrolPosition(grid [][]rune) (int, int) {
	y, x := -1, -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				y, x = i, j
				break
			}
		}
	}

	return y, x
}

func (d Day) getStep(direction int) (int, int) {
	direction %= 4
	switch direction {
	case 0:
		return -1, 0
	case 1:
		return 0, 1
	case 2:
		return 1, 0
	case 3:
		return 0, -1
	default:
		return -1, -1
	}
}

func (d Day) isOutOfBounds(i, j, n, m int) bool {
	return i < 0 || i >= n || j < 0 || j >= m
}

func (d Day) getSafePositions(grid [][]rune) (int, bool) {
	y, x := d.getPatrolPosition(grid)

	direction := 0

	visited := make([][][]bool, len(grid))
	for i := range visited {
		visited[i] = make([][]bool, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = make([]bool, 4)
		}
	}

	loop := false

	for !d.isOutOfBounds(y, x, len(grid), len(grid[0])) {

		if visited[y][x][direction%4] {
			loop = true
			break
		}

		visited[y][x][direction%4] = true

		dy, dx := d.getStep(direction)

		if !d.isOutOfBounds(y+dy, x+dx, len(grid), len(grid[0])) && grid[y+dy][x+dx] == '#' {
			direction++
		} else {
			y, x = y+dy, x+dx
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < 4; k++ {
				if visited[i][j][k] {
					res++
					break
				}
			}
		}
	}

	return res, loop
}

func (d Day) markSafePositions(grid [][]rune) [][]bool {
	y, x := d.getPatrolPosition(grid)

	direction := 0

	visited := make([][][]bool, len(grid))
	for i := range visited {
		visited[i] = make([][]bool, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = make([]bool, 4)
		}
	}

	for !d.isOutOfBounds(y, x, len(grid), len(grid[0])) {

		if visited[y][x][direction%4] {
			break
		}

		visited[y][x][direction%4] = true

		dy, dx := d.getStep(direction)

		if !d.isOutOfBounds(y+dy, x+dx, len(grid), len(grid[0])) && grid[y+dy][x+dx] == '#' {
			direction++
		} else {
			y, x = y+dy, x+dx
		}
	}

	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < 4; k++ {
				if visited[i][j][k] {
					vis[i][j] = true
					break
				}
			}
		}
	}

	return vis
}

func (d Day) readGrid() [][]rune {
	grid := strings.Split(d.ReadInput(), "\n")
	charGrid := make([][]rune, len(grid))
	for i, str := range grid {
		charGrid[i] = []rune(str)
	}
	return charGrid
}
