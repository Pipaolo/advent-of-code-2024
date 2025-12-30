package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"aoc2025/pkg/parser"
)

const day = 9

// Point represents a red tile coordinate
type Point struct {
	X, Y int
}

func main() {
	useExample := false // Use real input

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

// parsePoints converts input lines "x,y" into Point structs
func parsePoints(input *parser.Input) []Point {
	points := make([]Point, 0, len(input.Lines))
	for _, line := range input.Lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{X: x, Y: y})
	}
	return points
}

// abs returns the absolute value of n
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// findLargestRectangle finds the largest rectangle area using any two points as opposite corners
func findLargestRectangle(points []Point) int {
	// TODO(human): Implement this function
	// Given a slice of points (red tile coordinates), find the maximum rectangle area
	// where any two points serve as opposite corners of the rectangle.
	//
	// Hint: For two points (x1,y1) and (x2,y2), the rectangle area is:
	//       width × height = |x2-x1| × |y2-y1|
	//
	// Return the maximum area found (should be 50 for the example input)
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			width := abs(p2.X - p1.X)
			height := abs(p2.Y - p1.Y)
			area := (width + 1) * (height + 1)
			// Update maximum area if needed
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func solvePart1(input *parser.Input) any {
	points := parsePoints(input)
	return findLargestRectangle(points)
}

// Segment represents a horizontal or vertical line segment
type Segment struct {
	x1, y1, x2, y2 int
	isHorizontal   bool
}

// CompressedGrid uses coordinate compression to handle large coordinate ranges
type CompressedGrid struct {
	cells   [][]bool    // true = valid (red or green), false = invalid
	prefix  [][]int     // 2D prefix sum of invalid cells
	xCoords []int       // sorted unique X coordinates
	yCoords []int       // sorted unique Y coordinates
	xIndex  map[int]int // maps X coordinate to compressed index
	yIndex  map[int]int // maps Y coordinate to compressed index
}

// newCompressedGrid creates a coordinate-compressed grid from the polygon
func newCompressedGrid(points []Point) *CompressedGrid {
	// Collect all unique X and Y coordinates
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)

	for _, p := range points {
		xSet[p.X] = true
		ySet[p.Y] = true
	}

	// Also add coordinates from the line segments between consecutive points
	for i := 0; i < len(points); i++ {
		next := (i + 1) % len(points)
		p1, p2 := points[i], points[next]
		xSet[p1.X] = true
		xSet[p2.X] = true
		ySet[p1.Y] = true
		ySet[p2.Y] = true
	}

	// Sort coordinates
	xCoords := make([]int, 0, len(xSet))
	for x := range xSet {
		xCoords = append(xCoords, x)
	}
	sort.Ints(xCoords)

	yCoords := make([]int, 0, len(ySet))
	for y := range ySet {
		yCoords = append(yCoords, y)
	}
	sort.Ints(yCoords)

	// Create index maps
	xIndex := make(map[int]int)
	for i, x := range xCoords {
		xIndex[x] = i
	}
	yIndex := make(map[int]int)
	for i, y := range yCoords {
		yIndex[y] = i
	}

	// Create compressed cells grid
	height := len(yCoords)
	width := len(xCoords)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return &CompressedGrid{
		cells:   cells,
		xCoords: xCoords,
		yCoords: yCoords,
		xIndex:  xIndex,
		yIndex:  yIndex,
	}
}

// markPolygonInterior uses ray casting to determine which compressed cells are inside
func (g *CompressedGrid) markPolygonInterior(points []Point) {
	// Build list of segments from the polygon
	segments := make([]Segment, len(points))
	for i := 0; i < len(points); i++ {
		next := (i + 1) % len(points)
		p1, p2 := points[i], points[next]
		seg := Segment{x1: p1.X, y1: p1.Y, x2: p2.X, y2: p2.Y}
		seg.isHorizontal = (p1.Y == p2.Y)
		// Normalize so that x1 <= x2 and y1 <= y2
		if seg.x1 > seg.x2 {
			seg.x1, seg.x2 = seg.x2, seg.x1
		}
		if seg.y1 > seg.y2 {
			seg.y1, seg.y2 = seg.y2, seg.y1
		}
		segments[i] = seg
	}

	// For each compressed cell, check if it's inside the polygon
	for yi, y := range g.yCoords {
		for xi, x := range g.xCoords {
			if isPointInPolygon(x, y, segments) {
				g.cells[yi][xi] = true
			}
		}
	}
}

// isPointInPolygon uses ray casting (counting crossings to the right)
func isPointInPolygon(x, y int, segments []Segment) bool {
	crossings := 0

	for _, seg := range segments {
		if seg.isHorizontal {
			// Horizontal segment: check if point is ON the segment
			if y == seg.y1 && x >= seg.x1 && x <= seg.x2 {
				return true // On the boundary = inside
			}
		} else {
			// Vertical segment: check if ray from (x,y) going right crosses it
			if seg.x1 >= x && y >= seg.y1 && y < seg.y2 {
				crossings++
			}
			// Check if point is ON the vertical segment
			if x == seg.x1 && y >= seg.y1 && y <= seg.y2 {
				return true // On the boundary = inside
			}
		}
	}

	return crossings%2 == 1
}

// buildPrefixSum creates 2D prefix sum for O(1) rectangle queries
func (g *CompressedGrid) buildPrefixSum() {
	height := len(g.yCoords)
	width := len(g.xCoords)

	g.prefix = make([][]int, height+1)
	for i := range g.prefix {
		g.prefix[i] = make([]int, width+1)
	}

	for y := 1; y <= height; y++ {
		for x := 1; x <= width; x++ {
			invalid := 0
			if !g.cells[y-1][x-1] {
				invalid = 1
			}
			g.prefix[y][x] = invalid + g.prefix[y-1][x] + g.prefix[y][x-1] - g.prefix[y-1][x-1]
		}
	}
}

// isValidRectangle checks if rectangle from p1 to p2 contains only valid cells
func (g *CompressedGrid) isValidRectangle(p1, p2 Point) bool {
	// Get compressed indices
	xi1, ok1 := g.xIndex[p1.X]
	yi1, ok2 := g.yIndex[p1.Y]
	xi2, ok3 := g.xIndex[p2.X]
	yi2, ok4 := g.yIndex[p2.Y]

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return false
	}

	// Ensure proper ordering
	if xi1 > xi2 {
		xi1, xi2 = xi2, xi1
	}
	if yi1 > yi2 {
		yi1, yi2 = yi2, yi1
	}

	// Query prefix sum: any invalid cells in range?
	invalidCount := g.prefix[yi2+1][xi2+1] - g.prefix[yi1][xi2+1] - g.prefix[yi2+1][xi1] + g.prefix[yi1][xi1]
	return invalidCount == 0
}

// buildCompressedPolygonGrid creates the compressed grid with polygon interior marked
func buildCompressedPolygonGrid(points []Point) *CompressedGrid {
	grid := newCompressedGrid(points)
	grid.markPolygonInterior(points)
	grid.buildPrefixSum()
	return grid
}

func solvePart2(input *parser.Input) any {
	points := parsePoints(input)
	grid := buildCompressedPolygonGrid(points)

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if grid.isValidRectangle(points[i], points[j]) {
				width := abs(points[j].X - points[i].X)
				height := abs(points[j].Y - points[i].Y)
				area := (width + 1) * (height + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
