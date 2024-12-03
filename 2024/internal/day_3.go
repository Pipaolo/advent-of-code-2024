package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StartDay3Part1() {
	input, err := utilities.ParseFile("inputs/day_3_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Remove all line breaks and make the input a single line
	wholeInput := strings.Join(input, "")
	formattedInput := strings.ReplaceAll(strings.ReplaceAll(wholeInput, "\r\n", ""), "\n", "")
	expressionRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	digitsRe := regexp.MustCompile(`\d{1,3}`)
	matches := expressionRe.FindAllString(formattedInput, -1)

	total := 0
	for _, match := range matches {
		nums := digitsRe.FindAllString(match, -1)
		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])
		total += left * right
	}

	fmt.Println("Total Results:", total)
}

func StartDay3Part2() {
	input, err := utilities.ParseFile("inputs/day_3_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Remove all line breaks and make the input a single line
	wholeInput := strings.Join(input, "")
	formattedInput := strings.ReplaceAll(strings.ReplaceAll(wholeInput, "\r\n", ""), "\n", "")
	expressionRe := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	digitsRe := regexp.MustCompile(`\d{1,3}`)
	matches := expressionRe.FindAllString(formattedInput, -1)

	total := 0
	isDisabled := false
	for _, match := range matches {
		if match == "don't()" {
			isDisabled = true
		} else if match == "do()" {
			isDisabled = false
		} else {
			if !isDisabled {
				nums := digitsRe.FindAllString(match, -1)
				left, _ := strconv.Atoi(nums[0])
				right, _ := strconv.Atoi(nums[1])
				total += left * right
			}
		}
	}

	fmt.Println("Total Results:", total)
}
