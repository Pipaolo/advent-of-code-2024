# Day 6: Cephalopod Math - Visual Study Guide

## The Problem

We have a worksheet with numbers and operators. The last row contains operators (`+`, `-`, `*`, `/`).

```
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
```

---

## Part 1: Reading Horizontally

### How It Works

Split each row by spaces, then group tokens by their **position index**.

```
┌─────────────────────────────────────────────────────────────┐
│                    HORIZONTAL READING                        │
│                                                              │
│   Row 0:   "123"  "328"  "51"   "64"                        │
│              │      │      │      │                          │
│   Row 1:   "45"   "64"   "387"  "23"                        │
│              │      │      │      │                          │
│   Row 2:   "6"    "98"   "215"  "314"                       │
│              │      │      │      │                          │
│   Row 3:   "*"    "+"    "*"    "+"                         │
│              │      │      │      │                          │
│              ▼      ▼      ▼      ▼                          │
│           ┌─────┬─────┬─────┬─────┐                         │
│           │  0  │  1  │  2  │  3  │  ← Token Position       │
│           └─────┴─────┴─────┴─────┘                         │
└─────────────────────────────────────────────────────────────┘
```

### The Problems

```
Problem 0          Problem 1          Problem 2          Problem 3
─────────          ─────────          ─────────          ─────────
   123                328                 51                 64
    45                 64                387                 23
     6                 98                215                314
    *                  +                  *                  +
─────────          ─────────          ─────────          ─────────
123×45×6           328+64+98          51×387×215         64+23+314
= 33,210           = 490              = 4,245,255        = 401
```

### Algorithm (Part 1)

```
1. For each row:
   └── Split by whitespace → tokens[]
       └── For each token at position i:
           └── Add to problems[i]

2. Evaluate each problem with its operator
```

---

## Part 2: Reading Vertically

### The Key Insight

Now we read **each character column** as a separate number (top → bottom).

### Step 1: View as a Character Grid

```
     Column Index
     0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
     ┌─┬─┬─┬─┬─┬─┬─┬─┬─┬─┬──┬──┬──┬──┬──┐
  0  │1│2│3│ │3│2│8│ │ │5│1 │  │6 │4 │  │  ← Row 0
     ├─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼──┼──┼──┼──┼──┤
  1  │ │4│5│ │6│4│ │ │ │3│8 │7 │2 │3 │  │  ← Row 1
     ├─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼──┼──┼──┼──┼──┤
  2  │ │ │6│ │9│8│ │ │ │2│1 │5 │3 │1 │4 │  ← Row 2
     ├─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼──┼──┼──┼──┼──┤
  3  │*│ │ │ │+│ │ │ │ │*│  │  │+ │  │  │  ← Operator Row
     └─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴──┴──┴──┴──┴──┘
```

### Step 2: Find Problem Boundaries

Problems are separated by **columns that are ALL spaces**.

```
     0 1 2 │ 3 │ 4 5 6 │ 7 8 │ 9 10 11 │ 12 13 14
     ──────┼───┼───────┼─────┼─────────┼─────────
     1 2 3 │   │ 3 2 8 │     │ 5 1     │ 6  4
       4 5 │   │ 6 4   │     │ 3 8  7  │ 2  3
         6 │   │ 9 8   │     │ 2 1  5  │ 3  1  4
     *     │   │ +     │     │ *       │ +
     ──────┼───┼───────┼─────┼─────────┼─────────
           │ S │       │  S  │         │
           │ E │       │  E  │         │
           │ P │       │  P  │         │

     ├──────┤   ├───────┤     ├─────────┤─────────┤
      Prob 0     Prob 1         Prob 2    Prob 3
```

### Step 3: Read Each Column Vertically (Right → Left)

**Example: Problem 3 (columns 12-14)**

```
Reading RIGHT to LEFT within the problem:

     Col 14         Col 13         Col 12
     ──────         ──────         ──────
        ' '            '4'            '6'     ← Row 0
        ' '            '3'            '2'     ← Row 1
        '4'            '1'            '3'     ← Row 2
        ' '            ' '            '+'     ← Operator
         │              │              │
         ▼              ▼              ▼
     ┌───────┐     ┌───────┐     ┌───────┐
     │   4   │     │  431  │     │  623  │
     └───────┘     └───────┘     └───────┘

     First          Second         Third
     Number         Number         Number
```

**Result**: 4 + 431 + 623 = **1058**

### All Problems (Part 2)

```
┌─────────────────────────────────────────────────────────────┐
│ Problem 0 (cols 0-2)     │ Problem 1 (cols 4-6)            │
│                          │                                  │
│  Col: 2   1   0          │  Col: 6   5   4                  │
│       │   │   │          │       │   │   │                  │
│       3   2   1          │       8   2   3                  │
│       5   4              │           4   6                  │
│       6                  │           8   9                  │
│       ▼   ▼   ▼          │       ▼   ▼   ▼                  │
│      356  24   1         │       8  248  369                │
│                          │                                  │
│  356 × 24 × 1 = 8,544    │  8 + 248 + 369 = 625            │
├─────────────────────────────────────────────────────────────┤
│ Problem 2 (cols 8-10)    │ Problem 3 (cols 12-14)          │
│                          │                                  │
│  Col: 10  9   8          │  Col: 14  13  12                 │
│       │   │   │          │       │   │   │                  │
│       1   5              │       4   4   6                  │
│       8   3              │           3   2                  │
│       1   2              │           1   3                  │
│       ▼   ▼   ▼          │       ▼   ▼   ▼                  │
│      175 581  32         │       4  431  623                │
│                          │                                  │
│  175 × 581 × 32          │  4 + 431 + 623 = 1,058          │
│  = 3,253,600             │                                  │
└─────────────────────────────────────────────────────────────┘

Grand Total: 8,544 + 625 + 3,253,600 + 1,058 = 3,263,827
```

---

## Algorithm Summary

### Part 1: Horizontal
```
for each row:
    tokens = split by whitespace
    for i, token in tokens:
        problems[i].add(token)

for each problem:
    total += evaluate(numbers, operator)
```

### Part 2: Vertical
```
grid = convert to 2D character array

for col in range(width):
    if all rows are space at col:
        mark as separator

for each problem block (between separators):
    for col in block (RIGHT to LEFT):
        number = read digits TOP to BOTTOM
        add number to problem

    total += evaluate(numbers, operator)
```

---

## Key Data Structure: 2D Grid

```go
// Convert lines to uniform-width grid
func buildGrid(lines []string) [][]byte {
    maxLen := 0
    for _, line := range lines {
        if len(line) > maxLen {
            maxLen = len(line)
        }
    }

    grid := make([][]byte, len(lines))
    for i, line := range lines {
        grid[i] = make([]byte, maxLen)
        copy(grid[i], line)
        // Pad with spaces
        for j := len(line); j < maxLen; j++ {
            grid[i][j] = ' '
        }
    }
    return grid
}
```

---

## Complexity

| Part | Time | Space |
|------|------|-------|
| Part 1 | O(rows × tokens) | O(problems) |
| Part 2 | O(rows × cols) | O(rows × cols) |

---

## Answers

- **Part 1**: `4,722,948,564,882`
- **Part 2**: `9,581,313,737,063`
