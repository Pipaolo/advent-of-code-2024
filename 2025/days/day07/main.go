package main

import (
	"fmt"
	"log"

	"aoc2025/pkg/parser"
)

const day = 7

func main() {
	useExample := false // Use example for learning

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

type Position struct {
	x int
	y int
}

func (p *Position) IsEqual(other Position) bool {
	return p.x == other.x && p.y == other.y
}

type Splitter struct {
	Position Position
}

func (s *Splitter) HasBeamSides(t *TachyonManifold, l Position, r Position) (bool, bool) {
	return t.HasBeamAt(l), t.HasBeamAt(r)
}

func (s Splitter) String() string {
	return fmt.Sprintf("Splitter(Pos=%v)", s.Position)
}

type EmptySpace struct {
	Position Position
}

func (e EmptySpace) String() string {
	return fmt.Sprintf("EmptySpace(Pos=%v)", e.Position)
}

type Beam struct {
	Position Position
}

type TachyonManifold struct {
	StartingPosition Position
	Splitters        []Splitter
	SplitCounter     int
	EmptySpaces      []EmptySpace
	Beams            []Beam
	Width            int
	Height           int
}

func (t TachyonManifold) String() string {
	return fmt.Sprintf("TachyonManifold(Start=%d, Splitters=%v, EmptySpaces=%v)", t.StartingPosition, t.Splitters, t.EmptySpaces)
}

func (t TachyonManifold) HasBeamAt(pos Position) bool {
	for _, beam := range t.Beams {
		if beam.Position.IsEqual(pos) {
			return true
		}
	}
	return false
}

func (t TachyonManifold) HasSplitterAt(pos Position) bool {
	for _, splitter := range t.Splitters {
		if splitter.Position.IsEqual(pos) {
			return true
		}
	}
	return false
}

func (t TachyonManifold) HasEmptySpaceAt(pos Position) bool {
	for _, empty := range t.EmptySpaces {
		if empty.Position.IsEqual(pos) {
			return true
		}
	}
	return false
}

func (t *TachyonManifold) Traverse() {
	// Implement traversal logic here
	for i := 0; i < t.Height; i++ {
		for j := 0; j < t.Width; j++ {
			currentPos := Position{x: j, y: i}

			if currentPos.IsEqual(t.StartingPosition) {
				// Check the bottom of the starting position
				t.Beams = append(t.Beams, Beam{Position: Position{x: j, y: i + 1}})
				continue
			}

			if t.HasBeamAt(currentPos) {
				// Check if there is an empty space below
				downPos := Position{x: j, y: i + 1}

				if downPos.y >= t.Height {
					break // Out of bounds
				}

				if t.HasEmptySpaceAt(downPos) {
					t.Beams = append(t.Beams, Beam{Position: downPos})
					continue
				} else if t.HasSplitterAt(downPos) {
					// Handle splitter logic
					splitter := Splitter{Position: downPos}
					leftPos := Position{x: j - 1, y: i + 1}
					rightPos := Position{x: j + 1, y: i + 1}

					hasLeft, hasRight := splitter.HasBeamSides(t, leftPos, rightPos)

					if !hasLeft {
						t.Beams = append(t.Beams, Beam{Position: leftPos})
					}
					if !hasRight {
						t.Beams = append(t.Beams, Beam{Position: rightPos})
					}

					t.SplitCounter++
					continue
				}
			}
		}
	}
}

func solvePart1(input *parser.Input) any {
	// Build a tachyon manifold from the input
	manifold := &TachyonManifold{
		Splitters:   []Splitter{},
		EmptySpaces: []EmptySpace{},
		Beams:       []Beam{},
		Width:       len(input.Lines[0]),
		Height:      len(input.Lines),
	}

	for y, line := range input.Lines {
		// Parse the line and populate the manifold
		for x, char := range line {
			switch char {
			case 'S':
				manifold.StartingPosition = Position{x: x, y: y}
			case '^':
				manifold.Splitters = append(manifold.Splitters, Splitter{Position: Position{
					x: x, y: y,
				}})
			case '.':
				manifold.EmptySpaces = append(manifold.EmptySpaces, EmptySpace{Position: Position{
					x: x, y: y,
				}})
			}
		}
	}

	manifold.Traverse()

	return manifold.SplitCounter
}

func solvePart2(input *parser.Input) any {
	return solvePart2WithVisualization(input, false) // Set to true for debug trace
}

// =============================================================================
// PART 2: COUNTING TIMELINES (QUANTUM TACHYON MANIFOLD)
// =============================================================================
//
// CONCEPT: Think of it as a BINARY TREE where each splitter creates a branch:
//
//                          S (start)
//                          │
//                          ▼
//                          ^ (splitter 1)
//                        ╱   ╲
//                      L       R         ← 2 timelines after 1 splitter
//                      │       │
//                      ▼       ▼
//                      ^       ^         (splitters on each branch)
//                    ╱   ╲   ╱   ╲
//                   L    R  L    R       ← 4 timelines after 2 levels
//                   │    │  │    │
//                  ...  ... ... ...
//
// KEY INSIGHT: At each splitter, ONE timeline becomes TWO timelines.
//              The total timelines = sum of all "leaf" paths reaching the bottom.
//
// ALGORITHM: Recursive depth-first traversal
//   - Start from S, move down
//   - If we hit a splitter (^): recurse LEFT + recurse RIGHT, sum results
//   - If we reach bottom: return 1 (one complete timeline)
//
// =============================================================================

func solvePart2WithVisualization(input *parser.Input, debug bool) int {
	// Build a grid map for O(1) lookup instead of linear search
	grid := make(map[Position]rune)
	var startPos Position
	height := len(input.Lines)

	for y, line := range input.Lines {
		for x, char := range line {
			pos := Position{x: x, y: y}
			grid[pos] = char
			if char == 'S' {
				startPos = pos
			}
		}
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// MEMOIZATION CACHE - The key optimization!
	// ═══════════════════════════════════════════════════════════════════════════
	//
	// WHY DO WE NEED THIS?
	//
	// Without memoization, we recompute the same paths multiple times:
	//
	//        Path A              Path B
	//           \                  /
	//            \                /
	//             \              /
	//              →  (x=50)  ←      ← Both paths arrive at same position!
	//                  │
	//                  ▼
	//              (same sub-tree)
	//
	// Without cache: We compute sub-tree TWICE (exponential blowup!)
	// With cache: We compute sub-tree ONCE, reuse the result
	//
	// This turns O(2^n) into O(n) where n = number of unique positions
	//
	memo := make(map[Position]int)

	// countTimelines: Recursively count all paths from a position to the bottom
	// Now with memoization for O(n) complexity!
	var countTimelines func(pos Position, depth int) int
	countTimelines = func(pos Position, depth int) int {
		currentPos := pos

		for {
			// CHECK CACHE FIRST - Have we computed this before?
			if cached, ok := memo[currentPos]; ok {
				if debug {
					fmt.Printf("%sCache hit at (%d,%d) → %d timelines\n",
						indent(depth), currentPos.x, currentPos.y, cached)
				}
				return cached
			}

			// Move one step down
			nextPos := Position{x: currentPos.x, y: currentPos.y + 1}

			// BOTTOM REACHED - This path is ONE complete timeline!
			if nextPos.y >= height {
				if debug {
					fmt.Printf("%sReached bottom at x=%d → 1 timeline\n",
						indent(depth), currentPos.x)
				}
				memo[currentPos] = 1
				return 1
			}

			char := grid[nextPos]

			switch char {
			case '.':
				// Empty space - continue falling down
				currentPos = nextPos

			case '^':
				// ═══════════════════════════════════════════════════════════════
				// SPLITTER HIT! This is where the quantum magic happens!
				// ═══════════════════════════════════════════════════════════════
				leftPos := Position{x: nextPos.x - 1, y: nextPos.y}
				rightPos := Position{x: nextPos.x + 1, y: nextPos.y}

				if debug {
					fmt.Printf("%sSplitter at (%d,%d) → branching left(%d) and right(%d)\n",
						indent(depth), nextPos.x, nextPos.y, leftPos.x, rightPos.x)
				}

				// The magic formula: total = left_timelines + right_timelines
				leftTimelines := countTimelines(leftPos, depth+1)
				rightTimelines := countTimelines(rightPos, depth+1)

				total := leftTimelines + rightTimelines

				if debug {
					fmt.Printf("%s↳ Splitter (%d,%d) produced: %d + %d = %d timelines\n",
						indent(depth), nextPos.x, nextPos.y, leftTimelines, rightTimelines, total)
				}

				// CACHE THE RESULT before returning
				memo[currentPos] = total
				return total

			default:
				// Edge case: hit boundary or unexpected char
				memo[currentPos] = 1
				return 1
			}
		}
	}

	if debug {
		fmt.Println("\n=== TIMELINE COUNTING TRACE ===")
		fmt.Printf("Starting from S at (%d, %d)\n\n", startPos.x, startPos.y)
	}

	return countTimelines(startPos, 0)
}

// indent returns spacing for debug output tree visualization
func indent(depth int) string {
	result := ""
	for i := 0; i < depth; i++ {
		result += "  │ "
	}
	return result
}
