package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 3

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

	sum := 0
	for _, line := range input.Lines {
		joltages := []int{}
		for _, char := range strings.Split(strings.ReplaceAll(line, "\r", ""), "") {
			joltage, err := strconv.Atoi(char)

			if err != nil {
				log.Fatalf("Failed to convert character to integer: %v", err)
			}

			joltages = append(joltages, joltage)
		}

		highestJolstage := 0

		for i := 0; i < len(joltages); i++ {
			for j := i + 1; j < len(joltages); j++ {
				left := joltages[i]
				right := joltages[j]
				joltage, err := strconv.Atoi(fmt.Sprintf("%d%d", left, right))

				if err != nil {
					log.Fatalf("Failed to convert concatenated joltage to integer: %v", err)
				}

				if joltage > highestJolstage {
					highestJolstage = joltage
				}
			}
		}

		sum += highestJolstage
	}

	return sum
}

func solvePart2(input *parser.Input) any {
	var sum int64 = 0
	k := 12 // number of digits to select

	for _, line := range input.Lines {
		line = strings.ReplaceAll(line, "\r", "")
		digits := []int{}
		for _, char := range strings.Split(line, "") {
			digit, err := strconv.Atoi(char)
			if err != nil {
				log.Fatalf("Failed to convert character to integer: %v", err)
			}
			digits = append(digits, digit)
		}

		// Greedy selection: for each position, pick the largest digit
		// while ensuring enough digits remain for remaining positions
		n := len(digits)
		selected := []int{}
		startIdx := 0

		for i := 0; i < k; i++ {
			// We need (k - i - 1) more digits after this one
			// So we can pick from startIdx to n - (k - i - 1) - 1
			endIdx := n - (k - i - 1)
			maxDigit := -1
			maxIdx := startIdx

			for j := startIdx; j < endIdx; j++ {
				if digits[j] > maxDigit {
					maxDigit = digits[j]
					maxIdx = j
				}
			}

			selected = append(selected, maxDigit)
			startIdx = maxIdx + 1
		}

		// Convert selected digits to number
		numStr := ""
		for _, d := range selected {
			numStr += strconv.Itoa(d)
		}
		num, _ := strconv.ParseInt(numStr, 10, 64)
		sum += num
	}

	return sum
}
