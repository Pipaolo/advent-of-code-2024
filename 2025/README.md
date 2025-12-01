# Advent of Code 2025

Go project for Advent of Code 2025 solutions.

## Project Structure

```
.
├── days/           # Each day's solution
│   ├── day01/
│   │   └── main.go
│   ├── day02/
│   │   └── main.go
│   └── ...
├── inputs/         # Puzzle inputs
│   ├── day1.txt
│   ├── day1_example.txt
│   └── ...
├── pkg/
│   ├── parser/     # Universal input parser
│   └── utils/      # Common utilities
└── .vscode/        # Debug configurations
```

## Usage

### Running a Day

```bash
# From project root
go run ./days/day01

# Or with cd
cd days/day01 && go run .
```

### Debugging

1. Open the project in VS Code
2. Open a day's `main.go` file
3. Set breakpoints where needed
4. Use the Run and Debug panel (Ctrl+Shift+D / Cmd+Shift+D)
5. Select the appropriate day configuration or "Debug Current Day"
6. Press F5 to start debugging

### Creating a New Day

```bash
./newday.sh 2   # Creates days/day02/main.go and input files
```

## Parser Features

```go
input, _ := parser.ReadInput(1)      // Read day 1 input
input, _ := parser.ReadExample(1)    // Read day 1 example

// Available methods:
input.Lines                          // []string of all lines
input.Raw                            // Raw string content
input.ToInts()                       // Parse as []int
input.ToIntGrid()                    // Parse as [][]int (space-separated)
input.ToCharGrid()                   // Parse as [][]rune
input.SplitByEmptyLine()             // Group lines by empty lines
input.ParseWithDelimiter(",")        // Split each line by delimiter
```

## Utilities

```go
utils.Abs(x)                  // Absolute value
utils.Min(a, b)               // Minimum
utils.Max(a, b)               // Maximum
utils.Sum(nums)               // Sum of slice
utils.GCD(a, b)               // Greatest common divisor
utils.LCM(a, b)               // Least common multiple
utils.MustAtoi(s)             // String to int (panics on error)

// 2D Grid helpers
utils.Point2D{X: 0, Y: 0}     // 2D point
utils.Cardinals                // Up, Down, Left, Right directions
utils.InBounds(p, w, h)       // Check if point is in bounds
```
