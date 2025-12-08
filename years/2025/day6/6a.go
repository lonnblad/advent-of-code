package day6

import (
	"regexp"
	"strings"

	"github.com/lonnblad/advent-of-code/util/strconv"
)

func Day6A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")
	operators := parseOperators(rows[len(rows)-1])
	dataRows := rows[:len(rows)-1]

	columnResults := processColumns(dataRows, operators)

	var totalSum int
	for _, result := range columnResults {
		totalSum += result
	}

	return totalSum, nil
}

func parseOperators(lastRow string) []string {
	re := regexp.MustCompile(`([\+\*])`)
	matches := re.FindAllStringSubmatch(lastRow, -1)

	operators := make([]string, len(matches))
	for idx, match := range matches {
		operators[idx] = match[1]
	}

	return operators
}

func processColumns(dataRows []string, operators []string) []int {
	re := regexp.MustCompile(`(\d+)`)
	columnResults := make([]int, len(operators))

	for _, row := range dataRows {
		matches := re.FindAllStringSubmatch(row, -1)

		for idx, match := range matches {
			val := strconv.MustParseInt(match[1])

			switch operators[idx] {
			case "+":
				columnResults[idx] += val
			case "*":
				// For multiplication, initialize with first value if result is 0
				if columnResults[idx] == 0 {
					columnResults[idx] = val
				} else {
					columnResults[idx] *= val
				}
			}
		}
	}

	return columnResults
}
