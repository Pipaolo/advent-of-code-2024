package days

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

type PageOrder struct {
	Left  int
	Right int
}

func (p *PageOrder) String() string {
	return fmt.Sprintf("(%d, %d)", p.Left, p.Right)
}

func (p *PageOrder) IsValid(update PageUpdate) bool {

	leftIndex := -1
	rightIndex := -1

	for i, page := range update.Pages {
		if page == p.Left {
			leftIndex = i
		}
		if page == p.Right {
			rightIndex = i
		}

		if leftIndex != -1 && rightIndex != -1 {
			break
		}
	}

	if leftIndex == -1 || rightIndex == -1 {
		return false
	}

	return leftIndex < rightIndex
}

type PageUpdate struct {
	Pages []int
}

func parseInput(input []string) ([]PageOrder, []PageUpdate) {
	pageOrders := []PageOrder{}
	pageUpdates := []PageUpdate{}
	for _, line := range input {
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "|") {
			pageOrderSplit := strings.Split(trimmedLine, "|")
			left, _ := strconv.Atoi(pageOrderSplit[0])
			right, _ := strconv.Atoi(pageOrderSplit[1])
			pageOrders = append(pageOrders, PageOrder{Left: left, Right: right})
		} else {
			if trimmedLine == "" {
				continue
			}

			pageUpdateSplit := strings.Split(trimmedLine, ",")
			pageUpdate := PageUpdate{Pages: []int{}}
			for _, page := range pageUpdateSplit {
				pageInt, _ := strconv.Atoi(strings.TrimSpace(page))
				pageUpdate.Pages = append(pageUpdate.Pages, pageInt)
			}
			pageUpdates = append(pageUpdates, pageUpdate)
		}
	}

	return pageOrders, pageUpdates
}

func StartDay5Part1() {
	input, err := utilities.ParseFile("inputs/day_5_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	pageOrders, pageUpdates := parseInput(input)

	middleSum := 0
	for i, update := range pageUpdates {
		isValid := true

		// For each pair of pages in the update
		for i := 0; i < len(update.Pages); i++ {
			for j := i + 1; j < len(update.Pages); j++ {
				page1 := update.Pages[i]
				page2 := update.Pages[j]

				// Check if there's a rule requiring page2 to be before page1
				for _, rule := range pageOrders {
					if rule.Left == page2 && rule.Right == page1 {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			middleIndex := len(update.Pages) / 2
			middleSum += update.Pages[middleIndex]
			fmt.Printf("Update %d is valid: %v (middle page: %d)\n",
				i+1, update.Pages, update.Pages[middleIndex])
		} else {
			fmt.Printf("Update %d is invalid: %v\n", i+1, update.Pages)
		}
	}

	fmt.Println("Middle Sum: ", middleSum)
}

func StartDay5Part2() {
	input, err := utilities.ParseFile("inputs/day_5_input.txt")
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	pageOrders, pageUpdates := parseInput(input)
	invalidUpdates := []PageUpdate{}
	for i, update := range pageUpdates {
		isValid := true
		// For each pair of pages in the update
		for i := 0; i < len(update.Pages); i++ {
			for j := i + 1; j < len(update.Pages); j++ {
				page1 := update.Pages[i]
				page2 := update.Pages[j]

				// Check if there's a rule requiring page2 to be before page1
				for _, rule := range pageOrders {
					if rule.Left == page2 && rule.Right == page1 {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
			if !isValid {
				break
			}
		}

		if !isValid {
			fmt.Printf("Update %d is invalid: %v\n", i+1, update.Pages)
			invalidUpdates = append(invalidUpdates, update)
		}

	}

	middleSum := 0

	// Start fixing the invalid updates
	for i, invalidUpdate := range invalidUpdates {
		// Create a copy of the page to modify
		pages := make([]int, len(invalidUpdate.Pages))
		copy(pages, invalidUpdate.Pages)

		// Try to fix the update by swapping pages
		for i := 0; i < len(pages); i++ {
			for j := i + 1; j < len(pages); j++ {
				page1 := pages[i]
				page2 := pages[j]

				// Check if there's a rule requiring page2 to be before page1
				for _, rule := range pageOrders {
					if rule.Left == page2 && rule.Right == page1 {
						pages[i], pages[j] = pages[j], pages[i]
						break
					}
				}
			}
		}
		middleIndex := len(pages) / 2
		middleSum += pages[middleIndex]
		fmt.Printf("Update %d is valid: %v (middle page: %d)\n",
			i+1, pages, pages[middleIndex])
	}

	fmt.Println("Middle Sum: ", middleSum)
}
