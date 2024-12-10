package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strings"
)

const (
	Right = iota
	Down
	Left
	Up
	DiagonalDownRight
	DiagonalDownLeft
	DiagonalUpRight
	DiagonalUpLeft
)

type Direction struct {
	Direction int
	Row       int
	Col       int
}

func checkXmas(
	target string,
	grid [][]string,
	row int,
	col int,
	direction Direction,
) bool {
	lastRow := row + (len(target)-1)*direction.Row
	lastCol := col + (len(target)-1)*direction.Col

	if row < 0 || lastRow < 0 || row >= len(grid) || lastRow >= len(grid) ||
		col < 0 || lastCol < 0 || col >= len(grid[0]) || lastCol >= len(grid[0]) {
		return false
	}

	word := ""
	for i := 0; i < len(target); i++ {
		newRow := row + i*direction.Row
		newCol := col + i*direction.Col

		if grid[newRow][newCol] != string(target[i]) {
			return false
		}

		word += grid[newRow][newCol]
	}
	fmt.Println("Found word:", word)

	return true
}

func StartDay4Part1() {
	input, err := utilities.ParseFile("inputs/day_4_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	directions := []Direction{
		{Direction: Right, Row: 0, Col: 1},
		{Direction: Down, Row: 1, Col: 0},
		{Direction: Left, Row: 0, Col: -1},
		{Direction: Up, Row: -1, Col: 0},
		{Direction: DiagonalDownRight, Row: 1, Col: 1},
		{Direction: DiagonalDownLeft, Row: 1, Col: -1},
		{Direction: DiagonalUpRight, Row: -1, Col: 1},
		{Direction: DiagonalUpLeft, Row: -1, Col: -1},
	}
	// Make the input a 2d array that contains the letters
	wordSearchGrid := [][]string{}
	for _, line := range input {
		wordSearchGrid = append(wordSearchGrid, strings.Split(line, ""))
	}

	totalXmasCount := 0
	for rowIndex, row := range wordSearchGrid {
		for colIndex := range row {
			for _, direction := range directions {
				if checkXmas("XMAS", wordSearchGrid, rowIndex, colIndex, direction) {
					totalXmasCount++
				}
			}
		}
	}

	fmt.Println("Total Xmas Count:", totalXmasCount)
}

func StartDay4Part2() {
	input, err := utilities.ParseFile("inputs/day_4_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Make the input a 2d array that contains the letters
	wordSearchGrid := [][]string{}
	for _, line := range input {
		wordSearchGrid = append(wordSearchGrid, strings.Split(line, ""))
	}

	totalXCount := 0
	for rowIndex, row := range wordSearchGrid {
		for colIndex := range row {
			// Only check positions that contain 'A' as this must be the center of our X
			if wordSearchGrid[rowIndex][colIndex] != "A" {
				continue
			}

			// Skip if we're too close to any edge to form a complete X
			if rowIndex == 0 || rowIndex >= len(wordSearchGrid)-1 ||
				colIndex == 0 || colIndex >= len(row)-1 {
				continue
			}

			// Check both diagonals that form the X
			diagonal1 := ((wordSearchGrid[rowIndex-1][colIndex-1] == "M" && wordSearchGrid[rowIndex+1][colIndex+1] == "S") ||
				(wordSearchGrid[rowIndex-1][colIndex-1] == "S" && wordSearchGrid[rowIndex+1][colIndex+1] == "M"))

			diagonal2 := ((wordSearchGrid[rowIndex-1][colIndex+1] == "M" && wordSearchGrid[rowIndex+1][colIndex-1] == "S") ||
				(wordSearchGrid[rowIndex-1][colIndex+1] == "S" && wordSearchGrid[rowIndex+1][colIndex-1] == "M"))

			// Only count if both diagonals form valid MAS/SAM patterns
			if diagonal1 && diagonal2 {
				totalXCount++
			}
		}
	}

	fmt.Println("Total X-MAS Count:", totalXCount)
}
