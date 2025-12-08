package day7

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day7B(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)
	row, col := m.Find("S")
	startPos := [2]int{row, col}

	// Use a local memoization map to avoid global state issues
	memo := make(map[[2]int]int)

	return traverseB(m, startPos, memo), nil
}

// traverseB recursively traverses the matrix from the current position downward.
// Uses memoization to cache results and count the number of paths from each position.
// Returns the total number of paths from the start position to the bottom.
func traverseB(m matrix.Matrix, currentPos [2]int, memo map[[2]int]int) int {
	nextPos := [2]int{currentPos[0] + 1, currentPos[1]}

	// If we've reached the bottom, return 1 path
	if !m.IsWithinBounds(nextPos) {
		return 1
	}

	// Return cached result if available
	if cached, exists := memo[nextPos]; exists {
		return cached
	}

	var result int

	switch m.Values[nextPos[0]][nextPos[1]] {
	case ".":
		// Continue straight down - one path
		result = traverseB(m, nextPos, memo)

	case "^":
		// Split left and right - sum of paths from both directions
		leftPos := [2]int{nextPos[0], nextPos[1] - 1}
		rightPos := [2]int{nextPos[0], nextPos[1] + 1}

		result = traverseB(m, leftPos, memo) + traverseB(m, rightPos, memo)
	default:
		result = 0
	}

	// Cache the result before returning
	memo[nextPos] = result
	return result
}
