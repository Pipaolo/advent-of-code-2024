package main

import (
	"fmt"
	"log"

	"aoc2025/pkg/parser"
)

const day = 8

func main() {
	useExample := false // Use example for learning

	var input *parser.Input
	var err error

	if useExample {
		input, err = parser.ReadExample(day)
	} else {
		input, err = parser.ReadInput(day)
	}

	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Printf("=== Day %d ===\n", day)
	fmt.Printf("Part 1: %v\n", solvePart1(input))
	fmt.Println()
	fmt.Printf("Part 2: %v\n", solvePart2(input))
}

func solvePart1(input *parser.Input) any {
	return "not implemented"
}

func solvePart2(input *parser.Input) any {
	return "not implemented"
}
