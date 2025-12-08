package day1

import (
	"math"
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

		rotations := int(steps / 100)
		zeroPos += rotations
		steps = int(math.Mod(float64(steps), 100))

		if matches[1] == "L" {
			nextPos -= steps
			if nextPos < 0 {
				nextPos = 100 + nextPos

				if prevPos != 0 && nextPos != 0 {
					zeroPos += 1
				}
			}
		} else {
			nextPos += steps
			if nextPos >= 100 {
				nextPos = nextPos - 100

				if prevPos != 0 && nextPos != 0 {
					zeroPos += 1
				}

			}
		}

		if nextPos == 0 {
			zeroPos += 1
		}
	}

	return zeroPos, nil
}
