package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Report struct {
	// Array of numbers representing the levels of the report.
	Levels []int
}

func (r *Report) IsSafe() bool {
	// Get the prev and current level

	increasingCount := 0
	decreasingCount := 0
	highestDiff := 0.0

	for i := range r.Levels {
		if i == 0 {
			continue
		}
		curr := r.Levels[i]
		prev := r.Levels[i-1]

		if curr > prev {
			increasingCount++
		} else if curr < prev {
			decreasingCount++
		} else if curr == prev {
			// This means there is at least one level that is not increasing or decreasing.
			return false
		}

		diff := math.Abs(float64(curr - prev))
		if diff > highestDiff {
			highestDiff = diff
		}
	}

	if highestDiff > 3 {
		return false
	}

	// This means there is at least one level that is not increasing or decreasing.
	if increasingCount != 0 && decreasingCount != 0 {
		return false
	}

	return true
}

func StartDay2() {
	input, err := utilities.ParseFile("inputs/day_2_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	safeReports := []Report{}

	// Parse the input
	for _, line := range input {
		levelsRaw := strings.Split(line, " ")
		levels := []int{}

		for _, rawLevel := range levelsRaw {
			level, _ := strconv.Atoi(rawLevel)
			levels = append(levels, level)
		}

		report := Report{Levels: levels}
		if report.IsSafe() {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Println(len(safeReports))
}
