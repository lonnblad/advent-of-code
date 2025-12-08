package day7

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day7B(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)

	startPos := [2]int{0, 0}

	for idx := range m.Values[0] {
		if m.Values[0][idx] == "S" {
			startPos[1] = idx
			break
		}
	}

	return traverseB(m, startPos), nil
}

var mapOfTraversedPositionsB = make(map[[2]int]int)

// traverse the matrix from the start position to the end position
// return the total sum of splits that need to be made
func traverseB(m matrix.Matrix, currentPos [2]int) int {
	nextPos := [2]int{currentPos[0] + 1, currentPos[1]}

	if !m.IsWithinBounds(nextPos) {
		return 1
	}

	if val, exists := mapOfTraversedPositionsB[nextPos]; exists {
		return val
	}

	switch m.Values[nextPos[0]][nextPos[1]] {
	case ".":
		return traverseB(m, nextPos)
	case "^":
		leftPos := [2]int{nextPos[0], nextPos[1] - 1}
		rightPos := [2]int{nextPos[0], nextPos[1] + 1}

		val := traverseB(m, leftPos) + traverseB(m, rightPos)
		mapOfTraversedPositionsB[nextPos] = val
		return val
	}

	return 0
}
