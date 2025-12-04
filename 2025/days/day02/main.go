package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 2

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
	invalidIds := []string{}

	for _, line := range input.Lines {
		// Get the ranges from each line separated by a comma
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			ids := strings.Split(r, "-")
			fId, err := strconv.Atoi(ids[0])

			if err != nil {
				log.Fatalf("Failed to convert first ID to integer: %v", err)
			}

			sId, err := strconv.Atoi(ids[1])

			if err != nil {
				log.Fatalf("Failed to convert second ID to integer: %v", err)
			}

			for id := fId; id <= sId; id++ {
				idStr := strconv.Itoa(id)
				left := idStr[:len(idStr)/2]
				right := idStr[len(idStr)/2:]

				if left == right {
					invalidIds = append(invalidIds, idStr)
				}
			}
		}
	}

	sum := 0

	for _, id := range invalidIds {
		idInt, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalf("Failed to convert ID to integer: %v", err)
		}

		sum += idInt
	}

	return sum
}

func solvePart2(input *parser.Input) any {
	invalidIds := []int{}

	for _, line := range input.Lines {
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			ids := strings.Split(r, "-")
			fId, err := strconv.Atoi(ids[0])
			if err != nil {
				log.Fatalf("Failed to convert first ID to integer: %v", err)
			}

			sId, err := strconv.Atoi(ids[1])
			if err != nil {
				log.Fatalf("Failed to convert second ID to integer: %v", err)
			}

			for id := fId; id <= sId; id++ {
				if isRepeatingPattern(strconv.Itoa(id)) {
					invalidIds = append(invalidIds, id)
				}
			}
		}
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}

	return sum
}

// isRepeatingPattern checks if a string is made of a pattern repeated at least twice
// Algorithm: Try all possible pattern lengths from 1 to len/2
// A pattern length must divide the total length evenly
func isRepeatingPattern(s string) bool {
	n := len(s)

	// Try each possible pattern length
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		// Optimization: pattern length must divide n evenly
		if n%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		isValid := true

		// Check if the pattern repeats throughout the string
		for i := patternLen; i < n; i += patternLen {
			if s[i:i+patternLen] != pattern {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}
