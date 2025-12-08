package day6

import (
	"regexp"
	"strconv"
	"strings"
)

func Day6A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	lastRow := rows[len(rows)-1]
	re := regexp.MustCompile(`([\+\*])`)
	matches := re.FindAllStringSubmatch(lastRow, -1)

	operators := make([]string, len(matches))
	for idx, match := range matches {
		operators[idx] = match[0]
	}

	sums := make([]int, len(operators))

	re = regexp.MustCompile(`(\d+)`)

	for _, row := range rows[:len(rows)-1] {
		matches := re.FindAllStringSubmatch(row, -1)

		for idx, match := range matches {
			val, _ := strconv.Atoi(match[0])

			switch operators[idx] {
			case "+":
				sums[idx] += val
			case "*":
				if sums[idx] == 0 {
					sums[idx] = val
				} else {
					sums[idx] *= val
				}
			}
		}
	}

	var totalSum int
	for _, sum := range sums {
		totalSum += sum
	}

	return totalSum, nil
}
