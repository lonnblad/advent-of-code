package day7

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day7A(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)

	startPos := [2]int{0, 0}

	for idx := range m.Values[0] {
		if m.Values[0][idx] == "S" {
			startPos[1] = idx
			break
		}
	}

	return traverseA(m, startPos, 0), nil
}

var mapOfTraversedPositions = make(map[[2]int]bool)

// traverse the matrix from the start position to the end position
// return the total sum of splits that need to be made
func traverseA(m matrix.Matrix, currentPos [2]int, totalSum int) int {
	nextPos := [2]int{currentPos[0] + 1, currentPos[1]}

	if !m.IsWithinBounds(nextPos) {
		return totalSum
	}

	if mapOfTraversedPositions[nextPos] {
		return totalSum
	}

	mapOfTraversedPositions[nextPos] = true

	switch m.Values[nextPos[0]][nextPos[1]] {
	case ".":
		return traverseA(m, nextPos, totalSum)
	case "^":
		leftPos := [2]int{nextPos[0], nextPos[1] - 1}
		rightPos := [2]int{nextPos[0], nextPos[1] + 1}

		return totalSum + 1 +
			traverseA(m, leftPos, 0) +
			traverseA(m, rightPos, 0)
	}

	return totalSum
}
