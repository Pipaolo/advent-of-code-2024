# Day 7: Quantum Tachyon Manifold - Visual Study Guide

## The Problem

A particle falls through a manifold from `S` (start) at the top to the bottom. When it hits a splitter (`^`), it splits into TWO particles - one goes left, one goes right.

```
.......S.......    ← Particle starts here
...............
.......^.......    ← Splitter: particle splits into LEFT and RIGHT
...............
......^.^......    ← Each path may hit more splitters
...............
.....^.^.^.....
...............
```

**Part 1**: Count the splitters
**Part 2**: Count the total number of **timelines** (unique paths from S to bottom)

---

## Core Concept: The Binary Tree

Each splitter creates a BRANCH in the decision tree:

```
                              S (start)
                              │
                              ▼
                         ┌────^────┐     ← Splitter 1
                         │         │
                       LEFT      RIGHT
                         │         │
                    ┌────^────┐   ...    ← Splitter 2 (on left path)
                    │         │
                  LEFT      RIGHT
                    │         │
                   ...       ...

    Each leaf of this tree = ONE TIMELINE
```

---

## Part 2: Counting Timelines

### The Key Insight

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                   │
│   At each SPLITTER:                                              │
│                                                                   │
│       1 timeline in → 2 timelines out                            │
│                                                                   │
│                │                                                  │
│                ▼                                                  │
│                ^                                                  │
│              ╱   ╲                                                │
│            ╱       ╲                                              │
│          L           R                                           │
│                                                                   │
│   Total timelines = LEFT timelines + RIGHT timelines             │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
```

### Formula

```
countTimelines(position):

    if reached_bottom:
        return 1                    ← One complete path = one timeline

    if hit_splitter:
        left  = countTimelines(left_position)
        right = countTimelines(right_position)
        return left + right         ← Sum both branches

    else:
        keep falling down...
```

---

## Step-by-Step Example

Using the example input:

```
.......S.......   row 0
...............   row 1
.......^.......   row 2   ← Splitter at (7,2)
...............   row 3
......^.^......   row 4   ← Splitters at (6,4) and (8,4)
...............   row 5
.....^.^.^.....   row 6   ← Splitters at (5,6), (7,6), (9,6)
...............   row 7
....^.^...^....   row 8
...............   row 9
...^.^...^.^...   row 10
...............   row 11
..^...^.....^..   row 12
...............   row 13
.^.^.^.^.^...^.   row 14
...............   row 15  ← BOTTOM (row 16 would be off-grid)
```

### Step 1: Start at S

```
Position: (7, 0)   ← Start at S
Action: Move down
```

### Step 2: Fall until we hit something

```
(7, 0) → (7, 1) → ...wait, what's at (7, 2)?

.......S.......
.......|.......     ← falling...
.......^.......     ← HIT SPLITTER!

It's a splitter! Time to branch.
```

### Step 3: Branch at splitter (7, 2)

```
                    (7, 0) S
                       │
                       ▼
                    (7, 2) ^
                     ╱   ╲
                   ╱       ╲
              (6, 2)      (8, 2)
                │            │
          countTimelines  countTimelines
                │            │
               ???          ???
```

Now we need to count timelines from BOTH branches:
- `countTimelines(6, 2)` - the LEFT branch
- `countTimelines(8, 2)` - the RIGHT branch

### Step 4: Trace the LEFT branch (6, 2)

```
From (6, 2), fall down...

(6, 2) → (6, 3) → (6, 4) ← What's here?

......^.^......   ← There's a ^ at position (6, 4)!

Another splitter! Branch again:
- Left: (5, 4)
- Right: (7, 4)
```

### Step 5: The recursion tree grows

```
                              (7, 0) S
                                 │
                              (7, 2) ^
                             ╱         ╲
                        (6, 2)        (8, 2)
                           │             │
                        (6, 4) ^      (8, 4) ^
                        ╱     ╲       ╱     ╲
                    (5,4)  (7,4)  (7,4)  (9,4)
                       │      │      │      │
                      ...    ...    ...    ...
```

### Step 6: Reach the bottom

Eventually, each path reaches row 15 (bottom):

```
Path example: S → left → left → left → ... → bottom

.......S.......
.......|.......
......|^.......    ← took LEFT at (7,2)
......|........
.....|^.^......    ← took LEFT at (6,4)
.....|.........
....|^.^.^.....    ← took LEFT at (5,6)
....|..........
...|^.^...^....    ← took LEFT at (4,8)
...|...........
..|^.^...^.^...
..|............
.|^...^.....^..
.|.............
|^.^.^.^.^...^.    ← reached bottom at x=0
|..............

This path = 1 timeline
```

### Step 7: Sum all paths

```
countTimelines(7, 2):           ← First splitter
│
├── LEFT: countTimelines(6, 2)
│   └── Returns: 25 timelines
│
├── RIGHT: countTimelines(8, 2)
│   └── Returns: 15 timelines
│
└── TOTAL: 25 + 15 = 40 timelines
```

---

## The Problem: Exponential Blowup

### Without Optimization

```
           S
           │
           ^₁
         ╱   ╲
        ^₂    ^₃
       ╱ ╲   ╱ ╲
      ^   ^  ^   ^     ← 4 paths
     ...  ... ... ...

At depth N: potentially 2^N paths to compute!

For the real input (70+ splitters deep): 2^70 = 1,180,591,620,717,411,303,424
                                         ↑ This would take FOREVER
```

### The Insight: Overlapping Subproblems

```
        Path A              Path B
           \                  /
            \                /
             \              /
              →  (x=50)  ←      ← Both paths arrive at SAME position!
                  │
                  ▼
              (same sub-tree)

Without cache: We compute (x=50)'s sub-tree TWICE
With cache:    We compute it ONCE, reuse the result
```

---

## The Solution: Memoization

### What is Memoization?

```
┌─────────────────────────────────────────────────────────────────┐
│  MEMOIZATION = "Remember answers we already computed"           │
│                                                                  │
│  memo = {}  // Empty cache                                      │
│                                                                  │
│  countTimelines(pos):                                           │
│      if pos in memo:                                            │
│          return memo[pos]      ← Already computed? Return it!   │
│                                                                  │
│      ... compute answer ...                                     │
│                                                                  │
│      memo[pos] = answer        ← Store for future use           │
│      return answer                                              │
└─────────────────────────────────────────────────────────────────┘
```

### Visual Example

```
First visit to (50, 20):
┌──────────────────────────────────────┐
│  pos = (50, 20)                      │
│  memo[(50,20)] = ???                 │
│                                      │
│  Computing...                        │
│  └── Recurse left: 100 timelines    │
│  └── Recurse right: 150 timelines   │
│  └── Total: 250                      │
│                                      │
│  memo[(50,20)] = 250  ← SAVE IT!    │
│  return 250                          │
└──────────────────────────────────────┘

Second visit to (50, 20):
┌──────────────────────────────────────┐
│  pos = (50, 20)                      │
│  memo[(50,20)] = 250  ← FOUND!      │
│                                      │
│  return 250  ← INSTANT! No recursion │
└──────────────────────────────────────┘
```

### Complexity Improvement

```
┌─────────────────────────────────────────────────────────────┐
│                                                               │
│  WITHOUT memoization:  O(2^n)  ← Exponential (unusable)      │
│                                                               │
│  WITH memoization:     O(n)    ← Linear (instant!)           │
│                                                               │
│  Where n = number of unique positions in the grid            │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

---

## Complete Algorithm

```
┌─────────────────────────────────────────────────────────────────┐
│                    QUANTUM TIMELINE COUNTER                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  INPUT:  Grid with 'S', '^', and '.'                            │
│  OUTPUT: Number of unique timelines (paths from S to bottom)    │
│                                                                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  STEP 1: Parse the grid                                         │
│          └── Find 'S' position                                  │
│          └── Build map of all characters by position            │
│                                                                  │
│  STEP 2: Initialize memoization cache                           │
│          └── memo = empty map                                   │
│                                                                  │
│  STEP 3: Call countTimelines(startPosition)                     │
│                                                                  │
│  STEP 4: countTimelines(pos) algorithm:                         │
│          │                                                       │
│          ├── CHECK CACHE: if pos in memo, return memo[pos]      │
│          │                                                       │
│          ├── LOOP: Move down one row at a time                  │
│          │   │                                                   │
│          │   ├── If BOTTOM: return 1, cache it                  │
│          │   │                                                   │
│          │   ├── If EMPTY (.): continue falling                 │
│          │   │                                                   │
│          │   └── If SPLITTER (^):                               │
│          │       │                                               │
│          │       ├── left  = countTimelines(x-1, y)             │
│          │       ├── right = countTimelines(x+1, y)             │
│          │       ├── total = left + right                       │
│          │       ├── memo[pos] = total                          │
│          │       └── return total                               │
│          │                                                       │
│          └── END LOOP                                           │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Code Implementation (Go)

```go
func countQuantumTimelines(grid map[Position]rune, start Position, height int) int {
    // Memoization cache: position → number of timelines from that point
    memo := make(map[Position]int)

    var countTimelines func(pos Position) int
    countTimelines = func(pos Position) int {
        currentPos := pos

        for {
            // Check cache first
            if cached, ok := memo[currentPos]; ok {
                return cached
            }

            // Move down
            nextPos := Position{x: currentPos.x, y: currentPos.y + 1}

            // Reached bottom?
            if nextPos.y >= height {
                memo[currentPos] = 1
                return 1
            }

            char := grid[nextPos]

            switch char {
            case '.':
                // Empty space - keep falling
                currentPos = nextPos

            case '^':
                // SPLITTER - branch into two timelines
                leftPos := Position{x: nextPos.x - 1, y: nextPos.y}
                rightPos := Position{x: nextPos.x + 1, y: nextPos.y}

                leftTimelines := countTimelines(leftPos)
                rightTimelines := countTimelines(rightPos)

                total := leftTimelines + rightTimelines
                memo[currentPos] = total
                return total

            default:
                memo[currentPos] = 1
                return 1
            }
        }
    }

    return countTimelines(start)
}
```

---

## Debug Trace (Example Input)

Running with debug enabled shows the recursion tree:

```
=== TIMELINE COUNTING TRACE ===
Starting from S at (7, 0)

Splitter at (7,2) → branching left(6) and right(8)
  │ Splitter at (6,4) → branching left(5) and right(7)
  │   │ Splitter at (5,6) → branching left(4) and right(6)
  │   │   │ ...
  │   │   │ ↳ Splitter (5,6) produced: 10 + 7 = 17 timelines
  │   │ Splitter at (7,6) → branching left(6) and right(8)
  │   │   │ ...
  │   │   │ ↳ Splitter (7,6) produced: 7 + 1 = 8 timelines
  │   │ ↳ Splitter (6,4) produced: 17 + 8 = 25 timelines
  │ Splitter at (8,4) → branching left(7) and right(9)
  │   │ ...
  │   │ ↳ Splitter (8,4) produced: 8 + 7 = 15 timelines
  │ ↳ Splitter (7,2) produced: 25 + 15 = 40 timelines

Part 2: 40
```

---

## Visual: Three Example Timelines

### Timeline 1: Always go LEFT

```
.......S.......
.......|.......
......|^.......    ← LEFT
......|........
.....|^.^......    ← LEFT
.....|.........
....|^.^.^.....    ← LEFT
....|..........
...|^.^...^....    ← LEFT
...|...........
..|^.^...^.^...
..|............
.|^...^.....^..
.|.............
|^.^.^.^.^...^.    ← Ends at x=0
|..............
```

### Timeline 2: Alternate LEFT-RIGHT

```
.......S.......
.......|.......
......|^.......    ← LEFT
......|........
......^|^......    ← RIGHT
.......|.......
.....^|^.^.....    ← LEFT
......|........
....^.^|..^....    ← RIGHT
.......|.......
...^.^.|.^.^...
.......|.......
..^...^|....^..
.......|.......
.^.^.^|^.^...^.
......|........
```

### Timeline 3: Same end, different path

```
.......S.......
.......|.......
......|^.......    ← LEFT
......|........
.....|^.^......    ← LEFT
.....|.........
....|^.^.^.....    ← LEFT
....|..........
....^|^...^....    ← RIGHT (different from Timeline 2!)
.....|.........
...^.^|..^.^...
......|........
..^..|^.....^..
.....|.........
.^.^.^|^.^...^.
......|........
```

All three are DIFFERENT timelines, even though some end at the same position!

---

## Key Takeaways

```
┌─────────────────────────────────────────────────────────────────┐
│  1. TREE STRUCTURE                                              │
│     └── Each splitter creates a binary branch                   │
│     └── Total timelines = sum of all leaf paths                 │
│                                                                  │
│  2. RECURSION                                                   │
│     └── Base case: bottom reached → return 1                    │
│     └── Recursive case: splitter → LEFT + RIGHT                 │
│                                                                  │
│  3. MEMOIZATION IS ESSENTIAL                                    │
│     └── Without it: O(2^n) - impossibly slow                    │
│     └── With it: O(n) - instant                                 │
│     └── Key insight: same position = same sub-problem           │
│                                                                  │
│  4. WHY IT WORKS                                                │
│     └── If two paths reach the same (x, y)...                   │
│     └── ...the timelines from that point are IDENTICAL          │
│     └── So we only need to compute it ONCE                      │
└─────────────────────────────────────────────────────────────────┘
```

---

## Complexity

| Aspect | Without Memo | With Memo |
|--------|--------------|-----------|
| Time   | O(2^n)       | O(n)      |
| Space  | O(n) stack   | O(n) cache + O(n) stack |

Where n = number of unique positions that can be reached

---

## Answers

- **Part 1**: `1,613` splitters
- **Part 2**: `48,021,610,271,997` timelines (~48 trillion!)

---

## Practice Exercises

1. **Trace by hand**: For the example input, trace the first 3 levels of the recursion tree.

2. **Why memoization works**: Draw two different paths that reach position (6, 4). Why do they produce the same number of sub-timelines?

3. **Modify the algorithm**: What if splitters could have 3 branches instead of 2? How would you change the formula?

4. **Debug mode**: Run the code with `debug=true` on the example input and match the output to your hand-traced tree.
