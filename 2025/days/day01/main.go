package main

import (
	"fmt"
	"log"

	"aoc2025/pkg/parser"
)

const day = 1

func main() {
	// Set to true to use example input for debugging
	useExample := false

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
	fmt.Printf("Part 2: %v\n", solvePart2(input))
}

func solvePart1(input *parser.Input) any {
	position := 50
	var password int

	for _, line := range input.Lines {
		var action rune
		var value int
		fmt.Sscanf(line, "%c%d", &action, &value)

		if action == 'L' {
			value = -value
		}
		position = (position + value%100 + 100) % 100

		if position == 0 {
			password++
		}
	}

	return password
}

func solvePart2(input *parser.Input) any {
	position := 50
	var password int

	for _, line := range input.Lines {
		var action rune
		var value int
		_, err := fmt.Sscanf(line, "%c%d", &action, &value)
		if err != nil {
			log.Fatalf("Failed to parse line '%s': %v", line, err)
		}

		// Count full rotations (each full 100-step rotation passes 0 once)
		password += value / 100
		remaining := value % 100
		startPos := position

		switch action {
		case 'L':
			position = ((position-remaining)%100 + 100) % 100
			// Cross 0 if we started > 0 and moved enough steps to reach/pass it
			if startPos > 0 && remaining >= startPos {
				password++
			}
		case 'R':
			position = (position + remaining) % 100
			// Cross 0 if position + remaining >= 100 (wrapped past 99→0)
			if startPos+remaining >= 100 {
				password++
			}
		}
	}

	return password
}
