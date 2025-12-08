package day3

import (
	"strconv"
	"strings"
)

func Day3A(input string) (_ int, err error) {
	banks := strings.Split(input, "\n")

	var totalSum int

	for _, bank := range banks {
		a, b := findHighestPair(bank)
		// fmt.Println(a, b)

		totalSum += a*10 + b
	}

	return totalSum, nil
}

func findHighestPair(bank string) (int, int) {
	cells := make([]int, len(bank))
	for idx, cell := range strings.Split(bank, "") {
		cells[idx], _ = strconv.Atoi(cell)
	}

	highestIdx := 0
	secondHighestIdx := len(cells) - 1

	for i := 0; i < len(cells)-1; i++ {
		if cells[i] > cells[highestIdx] {
			highestIdx = i
			continue
		}
	}

	for i := len(cells) - 1; i > highestIdx; i-- {
		if cells[i] > cells[secondHighestIdx] {
			secondHighestIdx = i
		}
	}

	return cells[highestIdx], cells[secondHighestIdx]
}
