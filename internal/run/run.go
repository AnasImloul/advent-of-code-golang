package run

import (
	year2023 "github.com/AnasImloul/advent-of-code-golang/internal/2023"
	year2024 "github.com/AnasImloul/advent-of-code-golang/internal/2024"
)

func Solution(year, day int, part string) any {
	switch year {
	case 2023:
		return year2023.Run(day, part)
	case 2024:
		return year2024.Run(day, part)
	default:
		return nil
	}
}
