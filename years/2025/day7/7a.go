package day7

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day7A(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)
	row, col := m.Find("S")
	startPos := [2]int{row, col}

	// Use a local visited map to avoid global state issues
	visited := make(map[[2]int]bool)

	return traverseA(m, startPos, 0, visited), nil
}

// traverseA recursively traverses the matrix from the current position downward.
// When encountering a "^", it splits left and right, counting 1 split.
// Returns the total number of splits needed to reach the bottom.
func traverseA(m matrix.Matrix, currentPos [2]int, totalSum int, visited map[[2]int]bool) int {
	nextPos := [2]int{currentPos[0] + 1, currentPos[1]}

	// Check if we've reached the bottom of the matrix
	if !m.IsWithinBounds(nextPos) {
		return totalSum
	}

	// Skip if we've already visited this position
	if visited[nextPos] {
		return totalSum
	}

	visited[nextPos] = true

	switch m.Values[nextPos[0]][nextPos[1]] {
	case ".":
		// Continue straight down
		return traverseA(m, nextPos, totalSum, visited)

	case "^":
		// Split left and right, counting 1 split
		leftPos := [2]int{nextPos[0], nextPos[1] - 1}
		rightPos := [2]int{nextPos[0], nextPos[1] + 1}

		return totalSum + 1 +
			traverseA(m, leftPos, 0, visited) +
			traverseA(m, rightPos, 0, visited)
	}

	return totalSum
}
