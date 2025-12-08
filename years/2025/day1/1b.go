package day1

import (
	"regexp"
	"strconv"
	"strings"
)

func Day1B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	const startPos = 50
	var zeroPos int

	nextPos := startPos
	re := regexp.MustCompile(`([LR])(\d+)`)

	for _, row := range rows {
		prevPos := nextPos

		matches := re.FindStringSubmatch(row)
		steps, _ := strconv.Atoi(matches[2])

		rotations := steps / 100
		zeroPos += rotations

		steps %= 100

		if matches[1] == "L" {
			nextPos = (nextPos - steps + 100) % 100
		} else {
			nextPos = (nextPos + steps) % 100
		}

		if nextPos == 0 {
			zeroPos++
		}

		// Count crossing zero during wrap-around
		// This happens when we wrap around and don't land on zero
		if prevPos != 0 && nextPos != 0 {
			if matches[1] == "L" && prevPos < steps {
				zeroPos++ // Crossed 0 going left
			} else if matches[1] == "R" && prevPos > 100-steps {
				zeroPos++ // Crossed 0 going right
			}
		}
	}

	return zeroPos, nil
}
