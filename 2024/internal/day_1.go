package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func StartDay1() {
	input, err := utilities.ParseFile("inputs/day_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	leftArray := []int{}
	rightArray := []int{}

	for _, line := range input {
		split := strings.Split(line, "   ")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		leftArray = append(leftArray, left)
		rightArray = append(rightArray, right)
	}

	// Sort the arrays
	sort.Ints(leftArray)
	sort.Ints(rightArray)

	totalDistance := 0

	for i := 0; i < len(leftArray); i++ {
		left := leftArray[i]
		right := rightArray[i]
		distance := math.Abs(float64(left - right))
		totalDistance += int(distance)
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
}
