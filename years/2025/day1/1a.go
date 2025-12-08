package day1

import (
	"regexp"
	"strconv"
	"strings"
)

func Day1A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	const startPos = 50
	var zeroPos int

	nextPos := startPos
	re := regexp.MustCompile(`([LR])(\d+)`)

	for _, row := range rows {
		matches := re.FindStringSubmatch(row)

		steps, _ := strconv.Atoi(matches[2])
		steps %= 100

		if matches[1] == "L" {
			nextPos = (nextPos - steps + 100) % 100
		} else {
			nextPos = (nextPos + steps) % 100
		}

		if nextPos == 0 {
			zeroPos++
		}
	}

	return zeroPos, nil
}
