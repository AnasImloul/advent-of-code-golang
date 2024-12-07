package day_04

import (
	"strings"
)

func (d Day) getAllThreeLetterSquares() <-chan []string {
	lines := strings.Split(d.ReadInput(), "\n")

	n := len(lines)
	m := len(lines[0])

	// Create a channel to stream words
	wordChannel := make(chan []string)

	go func() {
		defer close(wordChannel)

		// Iterate through each starting position in the grid
		for i := 0; i <= n-3; i++ {
			for j := 0; j <= m-3; j++ {
				var square []string
				square = append(square, lines[i][j:j+3])
				square = append(square, lines[i+1][j:j+3])
				square = append(square, lines[i+2][j:j+3])
				wordChannel <- square
			}
		}
	}()

	return wordChannel
}

func (d Day) secondPart() any {
	res := 0
	for square := range d.getAllThreeLetterSquares() {
		if square[1][1] != 'A' {
			continue
		}
		if (square[0][0] == 'M' && square[0][2] == 'M' && square[2][0] == 'S' && square[2][2] == 'S') ||
			(square[0][0] == 'S' && square[0][2] == 'S' && square[2][0] == 'M' && square[2][2] == 'M') {
			res++
			continue
		}
		if (square[0][0] == 'M' && square[2][0] == 'M' && square[0][2] == 'S' && square[2][2] == 'S') ||
			(square[0][0] == 'S' && square[2][0] == 'S' && square[0][2] == 'M' && square[2][2] == 'M') {
			res++
			continue
		}
	}
	return res
}
