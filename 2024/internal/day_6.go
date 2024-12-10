package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
)

type PointType string

const (
	PointTypeObject PointType = "#"
	PointTypeEmpty  PointType = "."
	PointTypeGuard  PointType = "^"
	PointTypeGuard2 PointType = "v"
	PointTypeGuard3 PointType = ">"
	PointTypeGuard4 PointType = "<"
)

type Point struct {
	X    int
	Y    int
	Type PointType
}

type Map struct {
	Points []Point
	Guard  Guard
}

func (m *Map) Find(x int, y int) Point {
	for _, point := range m.Points {
		if point.X == x && point.Y == y {
			return point
		}
	}
	return Point{}
}

type GuardDirection string

const (
	GuardDirectionUp    GuardDirection = "^"
	GuardDirectionRight GuardDirection = ">"
	GuardDirectionDown  GuardDirection = "v"
	GuardDirectionLeft  GuardDirection = "<"
)

type Guard struct {
	CurrentPosition  Point
	CurrentDirection GuardDirection
	GuardMap         *Map
}

func parseDay6Input(input []string) (guardMap Map) {
	guardCharacters := []string{"^", "v", ">", "<"}

	for y, line := range input {
		for x, char := range line {
			for _, guardCharacter := range guardCharacters {
				if string(char) == guardCharacter {
					guard := Guard{
						CurrentPosition:  Point{X: x, Y: y},
						CurrentDirection: GuardDirection(guardCharacter),
						GuardMap:         &guardMap,
					}
					guardMap.Guard = guard
				}
			}

			guardMap.Points = append(guardMap.Points, Point{X: x, Y: y, Type: PointType(char)})
		}
	}

	return guardMap
}

var moves = [][]int{
	{-1, 0}, // Up (matches GuardDirectionUp)
	{0, 1},  // Right (matches GuardDirectionRight)
	{1, 0},  // Down (matches GuardDirectionDown)
	{0, -1}, // Left (matches GuardDirectionLeft)
}

func StartDay6Part1() {
	input, err := utilities.ParseFile("inputs/day_6_input.txt")
	if err != nil {
		fmt.Println("Error parsing file: ", err)
		return
	}

	guardMap := parseDay6Input(input)
	guard := &guardMap.Guard

	// Use a map to track visited positions
	visited := make(map[string]bool)
	moveIdx := 0 // 0=up, 1=right, 2=down, 3=left

	// Add initial position
	key := fmt.Sprintf("%d,%d", guard.CurrentPosition.Y, guard.CurrentPosition.X)
	visited[key] = true

	// Main loop
	for guard.CurrentPosition.Y > 0 &&
		guard.CurrentPosition.Y < len(input)-1 &&
		guard.CurrentPosition.X > 0 &&
		guard.CurrentPosition.X < len(input[0])-1 {

		// Calculate next position
		nextY := guard.CurrentPosition.Y + moves[moveIdx][0]
		nextX := guard.CurrentPosition.X + moves[moveIdx][1]

		// Check if next position is blocked
		nextPoint := guard.GuardMap.Find(nextX, nextY)
		if nextPoint.Type == PointTypeObject {
			// Turn right
			moveIdx = (moveIdx + 1) % 4
			continue
		}

		// Move to next position
		guard.CurrentPosition = Point{X: nextX, Y: nextY}

		// Record new position if we haven't been here
		key = fmt.Sprintf("%d,%d", guard.CurrentPosition.Y, guard.CurrentPosition.X)
		visited[key] = true
	}

	fmt.Println(len(visited))
}
