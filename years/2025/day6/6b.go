package day6

import (
	"strconv"
	"strings"
)

func Day6B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	lastRow := rows[len(rows)-1]

	noOfOperators := 0
	operators := make([]string, len(lastRow))

	for idx, cell := range strings.Split(lastRow, "") {
		if idx == 0 {
			operators[idx] = cell
			noOfOperators += 1

			continue
		}

		if cell == "+" || cell == "*" {
			operators[idx] = cell
			noOfOperators += 1

			if idx != 0 {
				operators[idx-1] = ""
			}

			continue
		}

		operators[idx] = operators[idx-1]
	}

	totalSum := 0

	var sum int
	for cdx := range lastRow {
		if operators[cdx] == "" {
			// fmt.Println(cdx, sum)
			totalSum += sum
			sum = 0

			continue
		}

		var val string

		for _, row := range rows[:len(rows)-1] {
			val += strings.Split(row, "")[cdx]
		}

		number, _ := strconv.Atoi(strings.TrimSpace(val))

		// fmt.Println(number)
		// fmt.Println(operators[cdx])

		if operators[cdx] == "+" {
			sum += number
		} else {
			if sum == 0 {
				sum = number
			} else {
				sum *= number
			}
		}
	}

	totalSum += sum

	return totalSum, nil
}
