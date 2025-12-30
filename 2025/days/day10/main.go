package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 10

func main() {
	useExample := false // Run on real input

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

// Machine represents one machine's configuration for Part 1
type Machine struct {
	// Target is a bitmask where bit i is set if light i should be ON
	Target    int
	NumLights int
	// Buttons is a list of bitmasks; each bitmask shows which lights that button toggles
	Buttons []int
}

// MachinePart2 represents one machine's configuration for Part 2
type MachinePart2 struct {
	// Targets are the joltage requirements for each counter
	Targets []int
	// Buttons is a list where each button is a slice of counter indices it affects
	Buttons [][]int
}

// parseMachine parses a line like: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
func parseMachine(line string) Machine {
	// Extract indicator pattern [...]
	indicatorRe := regexp.MustCompile(`\[([.#]+)\]`)
	indicatorMatch := indicatorRe.FindStringSubmatch(line)
	pattern := indicatorMatch[1]

	// Convert pattern to target bitmask
	target := 0
	for i, ch := range pattern {
		if ch == '#' {
			target |= (1 << i)
		}
	}

	// Extract button groups (...)
	buttonRe := regexp.MustCompile(`\(([0-9,]+)\)`)
	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)

	buttons := make([]int, len(buttonMatches))
	for i, match := range buttonMatches {
		// Parse comma-separated indices
		indices := strings.Split(match[1], ",")
		buttonMask := 0
		for _, idx := range indices {
			n, _ := strconv.Atoi(idx)
			buttonMask |= (1 << n)
		}
		buttons[i] = buttonMask
	}

	return Machine{
		Target:    target,
		NumLights: len(pattern),
		Buttons:   buttons,
	}
}

// TODO(human): Implement findMinPresses - the core algorithm
// Given a target bitmask and a list of button bitmasks, find the minimum
// number of buttons to press (each 0 or 1 times) to reach the target from 0.
//
// Hint: Since pressing a button twice cancels out, this is equivalent to
// finding the smallest subset of buttons whose XOR equals the target.
func findMinPresses(target int, buttons []int) int {
	minPresses := len(buttons) + 1 // Start with impossible value

	// Fix 1: Parentheses around (1 << len(buttons)) to get correct bound
	for i := 0; i < (1 << len(buttons)); i++ {
		xorSum := 0
		pressCount := 0
		for j := 0; j < len(buttons); j++ {
			if (i & (1 << j)) != 0 {
				xorSum ^= buttons[j]
				pressCount++
			}
		}

		// Fix 2: Track minimum instead of returning first match
		if xorSum == target && pressCount < minPresses {
			minPresses = pressCount
		}
	}

	if minPresses > len(buttons) {
		return -1 // No solution found
	}
	return minPresses
}

func solvePart1(input *parser.Input) any {
	total := 0
	for _, line := range input.Lines {
		if line == "" {
			continue
		}
		machine := parseMachine(line)
		minPresses := findMinPresses(machine.Target, machine.Buttons)
		fmt.Printf("Machine with target=%d: min presses = %d\n", machine.Target, minPresses)
		total += minPresses
	}
	return total
}

// parseMachinePart2 parses joltage requirements and buttons for Part 2
func parseMachinePart2(line string) MachinePart2 {
	// Extract joltage requirements {...}
	joltageRe := regexp.MustCompile(`\{([0-9,]+)\}`)
	joltageMatch := joltageRe.FindStringSubmatch(line)
	joltageStrs := strings.Split(joltageMatch[1], ",")

	targets := make([]int, len(joltageStrs))
	for i, s := range joltageStrs {
		targets[i], _ = strconv.Atoi(s)
	}

	// Extract button groups (...) as slices of indices
	buttonRe := regexp.MustCompile(`\(([0-9,]+)\)`)
	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)

	buttons := make([][]int, len(buttonMatches))
	for i, match := range buttonMatches {
		indices := strings.Split(match[1], ",")
		buttons[i] = make([]int, len(indices))
		for j, idx := range indices {
			buttons[i][j], _ = strconv.Atoi(idx)
		}
	}

	return MachinePart2{
		Targets: targets,
		Buttons: buttons,
	}
}

// Fraction represents a rational number for exact arithmetic
type Fraction struct {
	num, den int
}

func newFrac(n int) Fraction {
	return Fraction{n, 1}
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (f Fraction) simplify() Fraction {
	if f.den == 0 {
		return f
	}
	g := gcd(f.num, f.den)
	f.num /= g
	f.den /= g
	if f.den < 0 {
		f.num = -f.num
		f.den = -f.den
	}
	return f
}

func (f Fraction) add(g Fraction) Fraction {
	return Fraction{f.num*g.den + g.num*f.den, f.den * g.den}.simplify()
}

func (f Fraction) sub(g Fraction) Fraction {
	return Fraction{f.num*g.den - g.num*f.den, f.den * g.den}.simplify()
}

func (f Fraction) mul(g Fraction) Fraction {
	return Fraction{f.num * g.num, f.den * g.den}.simplify()
}

func (f Fraction) div(g Fraction) Fraction {
	return Fraction{f.num * g.den, f.den * g.num}.simplify()
}

func (f Fraction) isZero() bool {
	return f.num == 0
}

func (f Fraction) toInt() (int, bool) {
	if f.den == 0 {
		return 0, false
	}
	f = f.simplify()
	if f.den != 1 && f.den != -1 {
		return 0, false
	}
	return f.num / f.den, true
}

// findMinPressesPart2 tries all subsets of buttons and solves the linear system for each
func findMinPressesPart2(targets []int, buttons [][]int) int {
	numButtons := len(buttons)
	minTotal := -1

	// Try all subsets of buttons (2^numButtons possibilities)
	for mask := 1; mask < (1 << numButtons); mask++ {
		// Get indices of buttons in this subset
		var subset []int
		for i := 0; i < numButtons; i++ {
			if (mask & (1 << i)) != 0 {
				subset = append(subset, i)
			}
		}

		// Try to solve using only these buttons
		result := solveLinearSystem(targets, buttons, subset)
		if result != -1 {
			if minTotal == -1 || result < minTotal {
				minTotal = result
			}
		}
	}

	return minTotal
}

// solveLinearSystem solves the system using only the buttons in subset
// Returns total presses or -1 if no valid solution
func solveLinearSystem(targets []int, buttons [][]int, subset []int) int {
	numCounters := len(targets)
	numVars := len(subset)

	if numVars == 0 {
		// Check if all targets are 0
		for _, t := range targets {
			if t != 0 {
				return -1
			}
		}
		return 0
	}

	// Build augmented matrix [A | targets]
	matrix := make([][]Fraction, numCounters)
	for i := range matrix {
		matrix[i] = make([]Fraction, numVars+1)
		for j := range matrix[i] {
			matrix[i][j] = newFrac(0)
		}
		matrix[i][numVars] = newFrac(targets[i])
	}

	// Fill the coefficient matrix
	for col, btnIdx := range subset {
		for _, counterIdx := range buttons[btnIdx] {
			matrix[counterIdx][col] = newFrac(1)
		}
	}

	// Gaussian elimination
	pivotRow := 0
	pivotCols := make([]int, 0)

	for col := 0; col < numVars && pivotRow < numCounters; col++ {
		// Find pivot
		found := -1
		for r := pivotRow; r < numCounters; r++ {
			if !matrix[r][col].isZero() {
				found = r
				break
			}
		}
		if found == -1 {
			continue // No pivot in this column
		}

		// Swap rows
		matrix[pivotRow], matrix[found] = matrix[found], matrix[pivotRow]
		pivotCols = append(pivotCols, col)

		// Scale pivot row
		pivot := matrix[pivotRow][col]
		for c := col; c <= numVars; c++ {
			matrix[pivotRow][c] = matrix[pivotRow][c].div(pivot)
		}

		// Eliminate column
		for r := 0; r < numCounters; r++ {
			if r != pivotRow && !matrix[r][col].isZero() {
				factor := matrix[r][col]
				for c := col; c <= numVars; c++ {
					matrix[r][c] = matrix[r][c].sub(factor.mul(matrix[pivotRow][c]))
				}
			}
		}
		pivotRow++
	}

	// Check for inconsistency
	for r := pivotRow; r < numCounters; r++ {
		if !matrix[r][numVars].isZero() {
			return -1 // No solution
		}
	}

	// Identify free variables (columns without pivots)
	pivotSet := make(map[int]bool)
	for _, col := range pivotCols {
		pivotSet[col] = true
	}
	var freeVars []int
	for col := 0; col < numVars; col++ {
		if !pivotSet[col] {
			freeVars = append(freeVars, col)
		}
	}

	// If no free variables, extract unique solution
	if len(freeVars) == 0 {
		solution := make([]int, numVars)
		for i, col := range pivotCols {
			v, ok := matrix[i][numVars].toInt()
			if !ok || v < 0 {
				return -1
			}
			solution[col] = v
		}
		total := 0
		for _, v := range solution {
			total += v
		}
		return total
	}

	// With free variables, we need to search over them
	// Build the relationship: pivotVar[i] = rhs[i] - sum(coef[i][j] * freeVar[j])
	// where i indexes pivot rows, j indexes free variables

	// For each pivot row, store coefficients for free variables
	type pivotInfo struct {
		col   int       // which variable this pivot determines
		rhs   Fraction  // constant term
		coefs []Fraction // coefficients for each free variable (negated from matrix)
	}
	pivots := make([]pivotInfo, len(pivotCols))
	for i, pivCol := range pivotCols {
		info := pivotInfo{
			col:   pivCol,
			rhs:   matrix[i][numVars],
			coefs: make([]Fraction, len(freeVars)),
		}
		for j, freeCol := range freeVars {
			// The relationship is: pivotVar = rhs - coef*freeVar
			info.coefs[j] = matrix[i][freeCol]
		}
		pivots[i] = info
	}

	// Find bounds for free variables
	// Each free variable must be >= 0
	// Each pivot variable must be >= 0: rhs - sum(coef*free) >= 0
	maxFree := 0
	for _, t := range targets {
		if t > maxFree {
			maxFree = t
		}
	}

	// Search over free variable values (limited search)
	minTotal := -1
	var searchFreeVars func(idx int, freeVals []int)
	searchFreeVars = func(idx int, freeVals []int) {
		if idx == len(freeVars) {
			// Compute all variable values
			solution := make([]int, numVars)
			for i, fv := range freeVars {
				solution[fv] = freeVals[i]
			}
			for _, piv := range pivots {
				val := piv.rhs
				for j, coef := range piv.coefs {
					val = val.sub(coef.mul(newFrac(freeVals[j])))
				}
				v, ok := val.toInt()
				if !ok || v < 0 {
					return // Invalid
				}
				solution[piv.col] = v
			}
			total := 0
			for _, v := range solution {
				total += v
			}
			if minTotal == -1 || total < minTotal {
				minTotal = total
			}
			return
		}

		// Try values for this free variable
		for v := 0; v <= maxFree; v++ {
			freeVals[idx] = v
			searchFreeVars(idx+1, freeVals)
		}
	}

	if len(freeVars) <= 3 { // Limit search depth
		searchFreeVars(0, make([]int, len(freeVars)))
	}

	return minTotal
}

func solvePart2(input *parser.Input) any {
	total := 0
	for _, line := range input.Lines {
		if line == "" {
			continue
		}
		machine := parseMachinePart2(line)
		minPresses := findMinPressesPart2(machine.Targets, machine.Buttons)
		fmt.Printf("Machine with targets=%v: min presses = %d\n", machine.Targets, minPresses)
		total += minPresses
	}
	return total
}
