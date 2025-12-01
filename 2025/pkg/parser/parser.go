package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Input holds the raw and parsed input data
type Input struct {
	Raw   string
	Lines []string
}

// ReadInput reads the input file for a given day
func ReadInput(day int) (*Input, error) {
	return ReadFile(dayFilePath(day))
}

// ReadExample reads the example input file for a given day
func ReadExample(day int) (*Input, error) {
	return ReadFile(exampleFilePath(day))
}

// ReadFile reads any file and returns an Input struct
func ReadFile(path string) (*Input, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := string(data)
	lines := strings.Split(strings.TrimRight(raw, "\n"), "\n")

	return &Input{
		Raw:   raw,
		Lines: lines,
	}, nil
}

func dayFilePath(day int) string {
	return "inputs/day" + strconv.Itoa(day) + ".txt"
}

func exampleFilePath(day int) string {
	return "inputs/day" + strconv.Itoa(day) + "_example.txt"
}

// ToInts converts lines to integers
func (i *Input) ToInts() ([]int, error) {
	result := make([]int, 0, len(i.Lines))
	for _, line := range i.Lines {
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}

// ToIntGrid parses input into a 2D grid of integers (space-separated)
func (i *Input) ToIntGrid() ([][]int, error) {
	result := make([][]int, 0, len(i.Lines))
	for _, line := range i.Lines {
		if line == "" {
			continue
		}
		row := []int{}
		for _, field := range strings.Fields(line) {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}
		result = append(result, row)
	}
	return result, nil
}

// ToCharGrid parses input into a 2D grid of characters
func (i *Input) ToCharGrid() [][]rune {
	result := make([][]rune, 0, len(i.Lines))
	for _, line := range i.Lines {
		if line == "" {
			continue
		}
		result = append(result, []rune(line))
	}
	return result
}

// SplitByEmptyLine splits input into groups separated by empty lines
func (i *Input) SplitByEmptyLine() [][]string {
	var result [][]string
	var current []string

	for _, line := range i.Lines {
		if line == "" {
			if len(current) > 0 {
				result = append(result, current)
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}

// ParseWithDelimiter splits each line by a custom delimiter
func (i *Input) ParseWithDelimiter(delim string) [][]string {
	result := make([][]string, 0, len(i.Lines))
	for _, line := range i.Lines {
		if line == "" {
			continue
		}
		result = append(result, strings.Split(line, delim))
	}
	return result
}

// ForEachLine iterates over each line with a callback
func (i *Input) ForEachLine(fn func(lineNum int, line string)) {
	for idx, line := range i.Lines {
		fn(idx, line)
	}
}

// StreamInput reads input line by line (for very large files)
func StreamInput(day int, fn func(line string) error) error {
	file, err := os.Open(dayFilePath(day))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := fn(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}
