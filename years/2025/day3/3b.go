package day3

import (
	"strings"

	utilstrconv "github.com/lonnblad/advent-of-code/util/strconv"
)

func Day3B(input string) (_ int, err error) {
	banks := strings.Split(input, "\n")
	var totalSum int

	for _, bank := range banks {
		val := findHighest12Numbers(bank)
		totalSum += val
	}

	return totalSum, nil
}

func findHighest12Numbers(bank string) int {
	cells := make([]int, len(bank))

	for idx, cell := range strings.Split(bank, "") {
		cells[idx] = utilstrconv.MustParseInt(cell)
	}

	const numDigits = 12
	highestIndexes := make([]int, numDigits)

	// Find first highest digit
	highestIndexes[0] = 0

	for i := 1; i <= len(cells)-numDigits; i++ {
		if cells[i] > cells[highestIndexes[0]] {
			highestIndexes[0] = i

			if cells[i] == 9 {
				break // Can't get higher than 9
			}
		}
	}

	// Find remaining digits, each must appear after the previous one
	for idx := 1; idx < numDigits; idx++ {
		startIdx := highestIndexes[idx-1] + 1
		endIdx := len(cells) - (numDigits - idx)
		highestIndexes[idx] = startIdx

		for jdx := startIdx + 1; jdx <= endIdx; jdx++ {
			if cells[jdx] > cells[highestIndexes[idx]] {
				highestIndexes[idx] = jdx

				if cells[jdx] == 9 {
					break // Can't get higher than 9
				}
			}
		}
	}

	// Combine digits into a number
	var result int

	for _, idx := range highestIndexes {
		result = result*10 + cells[idx]
	}

	return result
}
