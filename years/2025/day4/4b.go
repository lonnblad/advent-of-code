package day4

import (
	"github.com/lonnblad/advent-of-code/util/matrix"
)

func Day4B(input string) (_ int, err error) {
	m := matrix.NewMatrix(input)

	var totalSum int

	for {
		noOfRollsThatCanBeMoved := 0

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
					m.Values[ridx][cidx] = "."
					noOfRollsThatCanBeMoved += 1
				}
			}
		}

		totalSum += noOfRollsThatCanBeMoved

		if noOfRollsThatCanBeMoved == 0 {
			return totalSum, nil
		}
	}
}
