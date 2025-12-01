package utils

import (
	"fmt"
	"sort"
	"strconv"
)

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of a slice of integers
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// GCD returns the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

// SortInts sorts a slice of integers in place and returns it
func SortInts(nums []int) []int {
	sort.Ints(nums)
	return nums
}

// Reverse reverses a slice of integers in place
func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// MustAtoi converts string to int, panics on error
func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %q to int: %v", s, err))
	}
	return n
}

// Point2D represents a 2D coordinate
type Point2D struct {
	X, Y int
}

// Add returns the sum of two points
func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{X: p.X + other.X, Y: p.Y + other.Y}
}

// Manhattan returns the Manhattan distance between two points
func (p Point2D) Manhattan(other Point2D) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Directions for 2D grid movement
var (
	Up        = Point2D{0, -1}
	Down      = Point2D{0, 1}
	Left      = Point2D{-1, 0}
	Right     = Point2D{1, 0}
	Cardinals = []Point2D{Up, Down, Left, Right}
	Diagonals = []Point2D{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	AllDirs   = append(Cardinals, Diagonals...)
)

// InBounds checks if a point is within grid bounds
func InBounds(p Point2D, width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}
