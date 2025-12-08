package day4

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day4A(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)
	var totalSum int

	for ridx, row := range m.Values {
		for cidx, cell := range row {
			if cell != "@" {
				continue
			}

			if countAdjacentRolls(m, ridx, cidx) < 4 {
				totalSum++
			}
		}
	}

	return totalSum, nil
}

func countAdjacentRolls(m matrix.Matrix, ridx, cidx int) int {
	adjacentCells := m.GetAdjacentCells([2]int{ridx, cidx})
	count := 0

	for _, adjacentCell := range adjacentCells {
		if adjacentCell == "@" {
			count++
		}
	}

	return count
}
