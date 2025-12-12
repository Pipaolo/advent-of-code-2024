package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 6

// Set to true to see step-by-step debug output
const DEBUG = true

func main() {
	useExample := true // Use example for learning

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

// =============================================================================
// PART 1: HORIZONTAL READING
// =============================================================================
//
// Input looks like:
//   123 328  51 64
//    45 64  387 23
//     6 98  215 314
//   *   +   *   +
//
// We split each row by spaces and group by POSITION:
//   Position 0: [123, 45, 6, *]   → 123 * 45 * 6
//   Position 1: [328, 64, 98, +]  → 328 + 64 + 98
//   Position 2: [51, 387, 215, *] → 51 * 387 * 215
//   Position 3: [64, 23, 314, +]  → 64 + 23 + 314
//
// =============================================================================

func solvePart1(input *parser.Input) any {
	if DEBUG {
		fmt.Println("========== PART 1: HORIZONTAL READING ==========")
		fmt.Println()
	}

	// A problem has a list of numbers and one operator
	type Problem struct {
		numbers  []int
		operator byte
	}

	// Map from position index → problem
	// Position 0 is the first token in each row, position 1 is second, etc.
	problemsByPosition := make(map[int]*Problem)

	// STEP 1: Process each row
	for rowIndex, rowText := range input.Lines {
		if DEBUG {
			fmt.Printf("Row %d: %q\n", rowIndex, rowText)
		}

		// Split the row by whitespace
		// strings.Fields automatically handles multiple spaces
		tokens := strings.Fields(rowText)

		if DEBUG {
			fmt.Printf("  Tokens: %v\n", tokens)
		}

		// STEP 2: Assign each token to its position
		for positionIndex, token := range tokens {
			// Create problem if it doesn't exist yet
			if problemsByPosition[positionIndex] == nil {
				problemsByPosition[positionIndex] = &Problem{}
			}

			currentProblem := problemsByPosition[positionIndex]

			// Check if this token is an operator (+, -, *, /)
			isOperatorToken := len(token) == 1 && (token[0] == '+' || token[0] == '-' || token[0] == '*' || token[0] == '/')

			if isOperatorToken {
				currentProblem.operator = token[0]
				if DEBUG {
					fmt.Printf("  Position %d: Found operator '%c'\n", positionIndex, token[0])
				}
			} else {
				// It's a number
				number, _ := strconv.Atoi(token)
				currentProblem.numbers = append(currentProblem.numbers, number)
				if DEBUG {
					fmt.Printf("  Position %d: Found number %d\n", positionIndex, number)
				}
			}
		}
		if DEBUG {
			fmt.Println()
		}
	}

	// STEP 3: Evaluate each problem and sum results
	if DEBUG {
		fmt.Println("--- Evaluating Problems ---")
	}

	grandTotal := 0
	for position := 0; position < len(problemsByPosition); position++ {
		problem := problemsByPosition[position]
		result := evaluateWithDebug(problem.numbers, problem.operator, DEBUG)
		grandTotal += result

		if DEBUG {
			fmt.Printf("Problem %d: %v %c = %d\n", position, problem.numbers, problem.operator, result)
		}
	}

	if DEBUG {
		fmt.Printf("\nGrand Total: %d\n", grandTotal)
	}

	return grandTotal
}

// =============================================================================
// PART 2: VERTICAL READING
// =============================================================================
//
// Same input, but now we read COLUMNS instead of rows:
//
//   Column:  0 1 2   3   4 5 6   7 8   9 10 11  12 13 14
//            ─────────────────────────────────────────────
//   Row 0:   1 2 3       3 2 8         5 1      6  4
//   Row 1:     4 5       6 4           3 8  7   2  3
//   Row 2:       6       9 8           2 1  5   3  1  4
//   Row 3:   *           +             *        +
//            ─────────────────────────────────────────────
//            └───┘       └───┘         └─────┘  └──────┘
//            Prob0       Prob1         Prob 2   Prob 3
//
// Each COLUMN becomes a number (read top to bottom):
//   Column 12: '6','2','3' → 623
//   Column 13: '4','3','1' → 431
//   Column 14: ' ',' ','4' → 4
//
// =============================================================================

func solvePart2(input *parser.Input) any {
	if DEBUG {
		fmt.Println("========== PART 2: VERTICAL READING ==========")
		fmt.Println()
	}

	// STEP 1: Convert input to a 2D grid of characters
	// This makes it easy to access any [row][column]
	grid := buildGridWithDebug(input.Lines)
	if len(grid) == 0 {
		return 0
	}

	numberOfRows := len(grid)
	numberOfColumns := len(grid[0])
	operatorRowIndex := numberOfRows - 1 // Last row has operators

	if DEBUG {
		fmt.Println("Grid created:")
		fmt.Printf("  Rows: %d, Columns: %d\n", numberOfRows, numberOfColumns)
		fmt.Printf("  Operator row is row %d\n", operatorRowIndex)
		fmt.Println()
		printGrid(grid)
		fmt.Println()
	}

	// STEP 2: Find which columns are "blank" (all spaces)
	// These blank columns separate different problems
	isColumnBlank := func(columnIndex int) bool {
		for rowIndex := 0; rowIndex < numberOfRows; rowIndex++ {
			character := grid[rowIndex][columnIndex]
			if character != ' ' {
				return false // Found a non-space, so column is NOT blank
			}
		}
		return true // All spaces, column IS blank
	}

	// STEP 3: Process the grid column by column
	// We'll find contiguous groups of non-blank columns (these are problems)

	grandTotal := 0
	problemNumber := 0
	currentColumn := 0

	for currentColumn < numberOfColumns {

		// STEP 3a: Skip over blank columns (these are separators)
		for currentColumn < numberOfColumns && isColumnBlank(currentColumn) {
			currentColumn++
		}

		// If we've gone past the end, we're done
		if currentColumn >= numberOfColumns {
			break
		}

		// STEP 3b: We found the start of a problem block
		problemStartColumn := currentColumn

		// Find where this problem block ends (next blank column)
		for currentColumn < numberOfColumns && !isColumnBlank(currentColumn) {
			currentColumn++
		}
		problemEndColumn := currentColumn // This is one past the last column

		if DEBUG {
			fmt.Printf("--- Problem %d: columns %d to %d ---\n",
				problemNumber, problemStartColumn, problemEndColumn-1)
		}

		// STEP 3c: Extract the operator and numbers from this block
		var operatorForThisProblem byte
		var numbersForThisProblem []int

		// We read columns from RIGHT to LEFT (as per the problem description)
		for col := problemEndColumn - 1; col >= problemStartColumn; col-- {

			// Check if this column has an operator (on the operator row)
			characterOnOperatorRow := grid[operatorRowIndex][col]
			if characterOnOperatorRow == '+' || characterOnOperatorRow == '-' ||
				characterOnOperatorRow == '*' || characterOnOperatorRow == '/' {
				operatorForThisProblem = characterOnOperatorRow
				if DEBUG {
					fmt.Printf("  Column %d: Found operator '%c'\n", col, characterOnOperatorRow)
				}
			}

			// Read the vertical number from this column
			// Go through rows 0 to operatorRowIndex-1 (data rows only)
			digitsFromTopToBottom := ""
			for row := 0; row < operatorRowIndex; row++ {
				character := grid[row][col]
				isDigit := character >= '0' && character <= '9'
				if isDigit {
					digitsFromTopToBottom += string(character)
				}
			}

			// If we found any digits, convert to a number
			if digitsFromTopToBottom != "" {
				number, _ := strconv.Atoi(digitsFromTopToBottom)
				numbersForThisProblem = append(numbersForThisProblem, number)
				if DEBUG {
					fmt.Printf("  Column %d: Digits '%s' → Number %d\n",
						col, digitsFromTopToBottom, number)
				}
			} else {
				if DEBUG {
					fmt.Printf("  Column %d: No digits found\n", col)
				}
			}
		}

		// STEP 3d: Evaluate this problem
		if len(numbersForThisProblem) > 0 && operatorForThisProblem != 0 {
			result := evaluateWithDebug(numbersForThisProblem, operatorForThisProblem, DEBUG)
			grandTotal += result
			if DEBUG {
				fmt.Printf("  Result: %v %c = %d\n\n",
					numbersForThisProblem, operatorForThisProblem, result)
			}
		}

		problemNumber++
	}

	if DEBUG {
		fmt.Printf("Grand Total: %d\n", grandTotal)
	}

	return grandTotal
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// buildGridWithDebug converts input lines into a 2D byte grid with uniform width
// Each row is padded with spaces to make all rows the same length
func buildGridWithDebug(lines []string) [][]byte {
	// Find the longest line
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Create the grid
	grid := make([][]byte, len(lines))
	for rowIndex, line := range lines {
		// Create a row with maxLength characters
		grid[rowIndex] = make([]byte, maxLength)

		// Copy the line content
		copy(grid[rowIndex], line)

		// Fill the rest with spaces
		for col := len(line); col < maxLength; col++ {
			grid[rowIndex][col] = ' '
		}
	}

	return grid
}

// printGrid displays the grid with row and column indices
func printGrid(grid [][]byte) {
	if len(grid) == 0 {
		return
	}

	// Print column headers
	fmt.Print("     ")
	for col := 0; col < len(grid[0]); col++ {
		fmt.Printf("%2d", col%10)
	}
	fmt.Println()

	fmt.Print("     ")
	for col := 0; col < len(grid[0]); col++ {
		fmt.Print("──")
	}
	fmt.Println()

	// Print each row
	for row := 0; row < len(grid); row++ {
		fmt.Printf("%2d │ ", row)
		for col := 0; col < len(grid[row]); col++ {
			ch := grid[row][col]
			if ch == ' ' {
				fmt.Print(" ·") // Show spaces as dots for visibility
			} else {
				fmt.Printf(" %c", ch)
			}
		}
		fmt.Println()
	}
}

// evaluateWithDebug applies an operator to a list of numbers
func evaluateWithDebug(numbers []int, operator byte, debug bool) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		previousResult := result
		currentNumber := numbers[i]

		switch operator {
		case '+':
			result = result + currentNumber
		case '-':
			result = result - currentNumber
		case '*':
			result = result * currentNumber
		case '/':
			result = result / currentNumber
		}

		if debug {
			fmt.Printf("    %d %c %d = %d\n", previousResult, operator, currentNumber, result)
		}
	}

	return result
}
