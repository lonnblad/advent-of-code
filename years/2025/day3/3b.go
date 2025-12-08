package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Day3B(input string) (_ int, err error) {
	banks := strings.Split(input, "\n")

	var totalSum int

	for _, bank := range banks {
		val := findHighest12Numbers(bank)
		fmt.Println(val)

		totalSum += val
		// break
	}

	return totalSum, nil
}

func findHighest12Numbers(bank string) int {
	cells := make([]int, len(bank))
	for idx, cell := range strings.Split(bank, "") {
		cells[idx], _ = strconv.Atoi(cell)
	}

	highestIndexes := make([]int, 12)

	for i := 0; i <= len(cells)-len(highestIndexes); i++ {
		if cells[i] > cells[highestIndexes[0]] {
			highestIndexes[0] = i

			if cells[i] == 9 {
				break
			}
		}
	}

	// fmt.Println(highestIndexes[0])

	for idx := 1; idx < 12; idx++ {
		highestIndexes[idx] = highestIndexes[idx-1] + 1

		for jdx := highestIndexes[idx-1] + 2; jdx <= len(cells)-(len(highestIndexes)-idx); jdx++ {
			// fmt.Println(jdx, cells[jdx], highestIndexes[idx], cells[highestIndexes[idx]])

			if cells[jdx] > cells[highestIndexes[idx]] {
				highestIndexes[idx] = jdx

				if cells[jdx] == 9 {
					break
				}
			}
		}
	}

	var result int

	for _, idx := range highestIndexes {
		result = result*10 + cells[idx]
	}

	return result
}
