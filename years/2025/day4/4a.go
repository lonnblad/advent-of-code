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

			adjacentCells := m.GetAdjacentCells([2]int{ridx, cidx})

			noOfRolls := 0
			for _, adjacentCell := range adjacentCells {
				if adjacentCell == "@" {
					noOfRolls += 1
				}
			}

			if noOfRolls < 4 {
				totalSum += 1
			}
		}
	}

	return totalSum, nil
}
