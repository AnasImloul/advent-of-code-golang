package run

import (
	year2023 "github.com/AnasImloul/advent-of-code-golang/internal/2023"
	year2024 "github.com/AnasImloul/advent-of-code-golang/internal/2024"
	"log"
)

func Solution(year, day int, part string) {
	switch year {
	case 2023:
		year2023.Run(day, part)
	case 2024:
		year2024.Run(day, part)
	default:
		log.Fatalf("Year %d not implemented", year)
	}
}
