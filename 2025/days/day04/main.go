package main

import (
	"fmt"
	"log"

	"aoc2025/pkg/parser"
)

const day = 4

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
	PAPER_ROLL := "@"

	accessibleRolls := 0
	for i, line := range input.Lines {
		// We need to check a character and their adjacents top, bottom, left, right and all diagonals
		for j, char := range line {
			if string(char) != PAPER_ROLL {
				continue
			}

			adjacentRolls := calculateAdjacentRolls(input.Lines, i, j, PAPER_ROLL)

			if adjacentRolls < 4 {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func solvePart2(input *parser.Input) any {
	PAPER_ROLL := "@"

	totalRemovedRolls := 0
	for {
		removedRolls := 0
		for i, line := range input.Lines {
			// We need to check a character and their adjacents top, bottom, left, right and all diagonals
			for j, char := range line {
				if string(char) != PAPER_ROLL {
					continue
				}

				if isAccessible(input.Lines, i, j, PAPER_ROLL) {
					// Store it's position
					removedRolls++
					// Replace it with a "x" to mark it as counted
					line = line[:j] + "x" + line[j+1:]
					input.Lines[i] = line
				}
			}
		}

		if removedRolls == 0 {
			break
		}

		totalRemovedRolls += removedRolls
	}

	return totalRemovedRolls
}

func isAccessible(lines []string, row, col int, targetChar string) bool {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(lines) && newCol >= 0 && newCol < len(lines[newRow]) {
			if string(lines[newRow][newCol]) == targetChar {
				count++
			}
		}
	}

	return count < 4
}

func calculateAdjacentRolls(lines []string, row, col int, targetChar string) int {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(lines) && newCol >= 0 && newCol < len(lines[newRow]) {
			if string(lines[newRow][newCol]) == targetChar {
				count++
			}
		}
	}

	return count
}
