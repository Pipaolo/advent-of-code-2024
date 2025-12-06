package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 5

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
	ranges := []string{}
	ingredients := []string{}

	for _, line := range input.Lines {
		if line == "" || line == "\r" {
			continue
		}

		// Parse line into ranges and ingredients
		if strings.Contains(line, "-") {
			// This is a range
			ranges = append(ranges, line)
			continue
		}

		// This is an ingredient
		ingredients = append(ingredients, line)
	}

	totalFresh := 0

	for _, ingredient := range ingredients {

		isFresh, _ := isIngredientFresh(ingredient, ranges)

		if isFresh {
			totalFresh++
		}
	}

	return totalFresh
}

type Range struct {
	start int
	end   int
}

func solvePart2(input *parser.Input) any {
	ranges := []Range{}

	for _, line := range input.Lines {
		if line == "" || line == "\r" {
			continue
		}

		if !strings.Contains(line, "-") {
			continue // Skip ingredient IDs
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		ranges = append(ranges, Range{start: start, end: end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{}
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[len(merged)-1]

		// Check if current range overlaps or touches the last merged range
		if r.start <= last.end+1 {
			// Merge: extend the end if needed
			if r.end > last.end {
				last.end = r.end
			}
		} else {
			merged = append(merged, r)
		}
	}

	total := 0
	for _, r := range merged {
		total += r.end - r.start + 1 // +1 because ranges are inclusive
	}

	return total
}

func isIngredientFresh(ingredient string, ranges []string) (bool, []int) {
	for _, r := range ranges {
		// Split the ranges
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}

		start, err := strconv.Atoi(parts[0])

		if err != nil {
			continue
		}

		end, err := strconv.Atoi(parts[1])

		if err != nil {
			continue
		}

		value, err := strconv.Atoi(ingredient)

		if err != nil {
			continue
		}

		if value >= start && value <= end {
			return true, []int{
				start,
				end,
			} // Fresh if in ANY range
		}
	}

	return false, []int{} // Spoiled if not in any range
}
