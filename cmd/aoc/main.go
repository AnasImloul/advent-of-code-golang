package main

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/generate"
	"github.com/AnasImloul/advent-of-code-golang/internal/run"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: aoc <'generate' || 'solve'> ...")
	}

	command := os.Args[1]

	switch command {
	case "generate":
		if len(os.Args) < 4 {
			log.Fatal("Usage: aoc generate <year> <day>")
		}
		year, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
		day, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Invalid day: %v", err)
		}
		if err := generate.GenerateFiles(year, day); err != nil {
			log.Fatalf("Error generating files: %v", err)
		}
	case "run":
		if len(os.Args) < 5 {
			log.Fatal("Usage: aoc solve <year> <day> <'first' || 'second'>")
		}
		year, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
		day, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Invalid day: %v", err)
		}
		part := os.Args[4]
		if part != "first" && part != "second" {
			log.Fatalf("Invalid part type: %s", part)
		}
		run.Solution(year, day, part)
	default:
		log.Fatalf("Invalid command: %s", command)
	}
}
