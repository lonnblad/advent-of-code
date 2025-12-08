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

				if countAdjacentRolls(m, ridx, cidx) < 4 {
					m.Values[ridx][cidx] = "."
					noOfRollsThatCanBeMoved++
				}
			}
		}

		totalSum += noOfRollsThatCanBeMoved

		if noOfRollsThatCanBeMoved == 0 {
			return totalSum, nil
		}
	}
}
