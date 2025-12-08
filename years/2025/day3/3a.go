package day3

import (
	"strings"

	utilstrconv "github.com/lonnblad/advent-of-code/util/strconv"
)

func Day3A(input string) (_ int, err error) {
	banks := strings.Split(input, "\n")
	var totalSum int

	for _, bank := range banks {
		a, b := findHighestPair(bank)
		totalSum += a*10 + b
	}

	return totalSum, nil
}

func findHighestPair(bank string) (int, int) {
	cells := make([]int, len(bank))

	for idx, cell := range strings.Split(bank, "") {
		cells[idx] = utilstrconv.MustParseInt(cell)
	}

	// Find highest digit (excluding last element)
	highestIdx := 0

	for i := 0; i < len(cells)-1; i++ {
		if cells[i] > cells[highestIdx] {
			highestIdx = i
		}
	}

	// Find second highest digit that appears after the highest
	secondHighestIdx := len(cells) - 1

	for i := len(cells) - 1; i > highestIdx; i-- {
		if cells[i] > cells[secondHighestIdx] {
			secondHighestIdx = i
		}
	}

	return cells[highestIdx], cells[secondHighestIdx]
}
