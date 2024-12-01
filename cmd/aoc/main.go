package main

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/2023"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: aoc <year> <day> <'first' || 'second>>")
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid year: %v", err)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid day: %v", err)
	}

	part := os.Args[3]
	if part != "first" && part != "second" {
		log.Fatalf("Invalid part type: %s", part)
	}

	runSolution(year, day, part)
}

func runSolution(year, day int, part string) {
	switch year {
	case 2023:
		year2023.Run(day, part)
	case 2024:
		year2024.Run(day, part)
	default:
		log.Fatalf("Year %d not implemented", year)
	}
}
