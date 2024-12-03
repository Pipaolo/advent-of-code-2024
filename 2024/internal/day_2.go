package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

type Report struct {
	// Array of numbers representing the levels of the report.
	Levels []*Level
}

type Level struct {
	Value         int
	IsDecreasing  bool
	IsProblematic bool
	Diff          float64
}

func (r *Report) IsSafeWithDampener() bool {

	// If the report is safe without the dampener, then return true
	if r.IsSafe(){ 
		return true
	}

	// If the report is not safe, then we need to check if it is safe with a dampener
	for i := range r.Levels {
		if i == 0 {
			continue
		}

		curr := r.Levels[i]
		prev := r.Levels[i-1]

		curr.IsProblematic = curr.Value == prev.Value || curr.IsDecreasing != prev.IsDecreasing || curr.Diff > 3
	}

	// Check for the problematic levels and try to remove them then check if the report is safe
	isSafe := false

	for i, level := range r.Levels {
		// This tries to check every possible combination of levels
		if level.IsProblematic {
			// Create a new report without the problematic level
			newReport := Report{Levels: append(r.Levels[:i], r.Levels[i+1:]...)}
			if newReport.IsSafe() {
				isSafe = true
				break
			}
		}
	}

	return isSafe
}

func (r *Report) IsSafe() bool {

	for i := range r.Levels {
		// We don't need to check the first two levels
		// As by default they are not decreasing
		if i == 0 {
			continue
		}

		curr := r.Levels[i]
		prev := r.Levels[i-1]

		curr.IsDecreasing = curr.Value < prev.Value

		// To make sure that the diff is always positive
		if curr.Value < prev.Value {
			curr.Diff = float64(prev.Value - curr.Value)
		} else {
			curr.Diff = float64(curr.Value - prev.Value)
		}

		if i == 1 {
			// Set the previous level to the current level
			// As we don't need to check the first level
			prev.IsDecreasing = curr.IsDecreasing
			prev.Diff = curr.Diff
		}

		if curr.Diff > 3 {
			return false
		}

		if curr.IsDecreasing != prev.IsDecreasing {
			return false
		}
	}

	return true
}

func levelInputToLevel(input string) *Level {
	level, _ := strconv.Atoi(input)
	return &Level{Value: level, IsDecreasing: false, IsProblematic: false, Diff: 0}
}

func inputToReports(input []string) []Report {
	reports := []Report{}
	for _, line := range input {
		report := Report{}
		for _, level := range strings.Split(line, " ") {
			report.Levels = append(report.Levels, levelInputToLevel(level))
		}
		reports = append(reports, report)
	}
	return reports
}

func StartDay2Part1() {
	input, err := utilities.ParseFile("inputs/day_2_input-test.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	reports := inputToReports(input)
	safeReports := []Report{}

	// Parse the input
	for _, report := range reports {
		if report.IsSafe() {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Println("Safe reports:", len(safeReports))
}

func StartDay2Part2() {
	input, err := utilities.ParseFile("inputs/day_2_input-test.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	reports := inputToReports(input)
	safeReports := []Report{}

	for _, report := range reports {
		if report.IsSafeWithDampener() {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Println("Safe reports with dampener:", len(safeReports))
}
