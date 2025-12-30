package main

import (
	"aoc2025/pkg/parser"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

const day = 8

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

type JunctionBox struct {
	Pos Position
}

type Position struct {
	X int
	Y int
	Z int
}

// Compute distance using the straight line distance formula
func computeStraightLineDistance(a, b Position) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// parseJunctionBoxes parses input lines into junction box positions
func parseJunctionBoxes(input *parser.Input) []Position {
	var positions []Position
	for _, line := range input.Lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		positions = append(positions, Position{X: x, Y: y, Z: z})
	}
	return positions
}

// Pair represents two junction boxes and their distance
type Pair struct {
	I, J     int // indices of the two junction boxes
	Distance float64
}

// UnionFind tracks which junction boxes are in the same circuit
type UnionFind struct {
	parent []int // parent[i] = parent of node i (or itself if root)
	rank   []int // rank[i] = approximate depth of subtree rooted at i
	size   []int // size[i] = number of nodes in circuit rooted at i
}

// NewUnionFind creates a new Union-Find structure for n elements
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i // each element is its own parent (separate circuit)
		uf.size[i] = 1   // each circuit starts with size 1
	}
	return uf
}

// Find returns the root of the circuit containing element x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges the circuits containing elements x and y
// Returns true if they were in different circuits (and got merged)
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // already in same circuit
	}

	// Union by rank: attach smaller tree under larger tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
		uf.rank[rootX]++
	}
	return true
}

// GetSize returns the size of the circuit containing element x
func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

func solvePart1(input *parser.Input) any {
	positions := parseJunctionBoxes(input)
	n := len(positions)

	// Generate all pairs with their distances
	var pairs []Pair
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := computeStraightLineDistance(positions[i], positions[j])
			pairs = append(pairs, Pair{I: i, J: j, Distance: dist})
		}
	}

	// Sort pairs by distance (shortest first)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].Distance < pairs[b].Distance
	})

	// Create Union-Find and make connections
	uf := NewUnionFind(n)
	connections := 1000
	if len(pairs) < connections {
		connections = len(pairs)
	}

	for i := 0; i < connections; i++ {
		uf.Union(pairs[i].I, pairs[i].J)
	}

	// Find sizes of all unique circuits
	circuitSizes := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		circuitSizes[root] = uf.size[root]
	}

	// Get the three largest circuit sizes
	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// Multiply the three largest
	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return result
}

func solvePart2(input *parser.Input) any {
	positions := parseJunctionBoxes(input)
	n := len(positions)

	// Generate all pairs with their distances
	var pairs []Pair
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := computeStraightLineDistance(positions[i], positions[j])
			pairs = append(pairs, Pair{I: i, J: j, Distance: dist})
		}
	}

	// Sort pairs by distance (shortest first)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].Distance < pairs[b].Distance
	})

	// Create Union-Find and connect until all in one circuit
	uf := NewUnionFind(n)
	circuitCount := n // Start with n separate circuits

	var lastPair Pair
	for _, pair := range pairs {
		if uf.Union(pair.I, pair.J) {
			// Successfully merged two circuits
			circuitCount--
			lastPair = pair

			if circuitCount == 1 {
				// All boxes are now in one circuit!
				break
			}
		}
	}

	// Multiply X coordinates of the last two connected boxes
	return positions[lastPair.I].X * positions[lastPair.J].X
}
