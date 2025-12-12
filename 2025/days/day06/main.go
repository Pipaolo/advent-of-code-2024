package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 6

func main() {
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

// buildGrid converts input lines into a 2D byte grid with uniform width
func buildGrid(lines []string) [][]byte {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, maxLen)
		copy(grid[i], line)
		for j := len(line); j < maxLen; j++ {
			grid[i][j] = ' '
		}
	}
	return grid
}

// isOperator checks if a character is an operator
func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

// evaluate applies an operator to a list of numbers
func evaluate(numbers []int, op byte) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		switch op {
		case '+':
			result += numbers[i]
		case '-':
			result -= numbers[i]
		case '*':
			result *= numbers[i]
		case '/':
			result /= numbers[i]
		}
	}
	return result
}

// solvePart1: Read numbers horizontally, group by token position
func solvePart1(input *parser.Input) any {
	// Group numbers by their token index (column position after splitting by spaces)
	type problem struct {
		numbers []int
		op      byte
	}
	problems := make(map[int]*problem)

	for _, row := range input.Lines {
		tokens := strings.Fields(row) // Split by whitespace, removes empty strings
		for i, token := range tokens {
			if problems[i] == nil {
				problems[i] = &problem{}
			}
			if len(token) == 1 && isOperator(token[0]) {
				problems[i].op = token[0]
			} else {
				num, _ := strconv.Atoi(token)
				problems[i].numbers = append(problems[i].numbers, num)
			}
		}
	}

	total := 0
	for _, p := range problems {
		total += evaluate(p.numbers, p.op)
	}
	return total
}

// solvePart2: Read numbers vertically (each column is a number), group by problem blocks
func solvePart2(input *parser.Input) any {
	grid := buildGrid(input.Lines)
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	opRow := rows - 1 // Last row contains operators

	// Find problem boundaries: columns where ALL rows are spaces
	isBlank := func(col int) bool {
		for row := 0; row < rows; row++ {
			if grid[row][col] != ' ' {
				return false
			}
		}
		return true
	}

	total := 0

	// Process each problem block (contiguous non-blank columns)
	col := 0
	for col < cols {
		// Skip blank columns
		for col < cols && isBlank(col) {
			col++
		}
		if col >= cols {
			break
		}

		// Find the extent of this problem block
		startCol := col
		for col < cols && !isBlank(col) {
			col++
		}
		endCol := col

		// Extract operator and numbers from this block
		var op byte
		var numbers []int

		for c := endCol - 1; c >= startCol; c-- { // Right to left
			// Check for operator on the operator row
			if isOperator(grid[opRow][c]) {
				op = grid[opRow][c]
			}

			// Read vertical number from data rows (top to bottom)
			numStr := ""
			for r := 0; r < opRow; r++ {
				if grid[r][c] >= '0' && grid[r][c] <= '9' {
					numStr += string(grid[r][c])
				}
			}
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
		}

		if len(numbers) > 0 && op != 0 {
			total += evaluate(numbers, op)
		}
	}

	return total
}
