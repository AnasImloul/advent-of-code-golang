package day_04

import (
	"fmt"
	"strings"
)

func (d Day) getAllFourLetterWords() <-chan string {
	lines := strings.Split(d.ReadInput(), "\n")

	n := len(lines)
	m := len(lines[0])

	// Define offsets for 8 directions
	directions := []struct{ dx, dy int }{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}

	// Helper function to check if a position is valid
	isValid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}

	// Create a channel to stream words
	wordChannel := make(chan string)

	go func() {
		defer close(wordChannel)

		// Iterate through each starting position in the grid
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				// For each direction
				for _, dir := range directions {
					// Collect the word for this direction
					word := ""
					for step := 0; step < 4; step++ {
						x := i + step*dir.dx
						y := j + step*dir.dy
						if !isValid(x, y) {
							break
						}
						word += string(lines[x][y])
					}
					// If a valid 4-letter word is formed, send it to the channel
					if len(word) == 4 {
						wordChannel <- word
					}
				}
			}
		}
	}()

	return wordChannel
}

func (d Day) firstPart() {
	res := 0
	for word := range d.getAllFourLetterWords() {
		if word == "XMAS" {
			res++
		}
	}
	fmt.Println(res)
}
