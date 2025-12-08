package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func Day2B(input string) (_ int, err error) {
	intervals := strings.Split(input, ",")

	var totalSum int

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, interval := range intervals {
		matches := re.FindStringSubmatch(interval)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])

		for i := start; i <= end; i++ {
			val := strconv.Itoa(i)

			for positions := 1; positions <= len(val)/2; positions++ {
				if hasSequence(val, positions) {
					totalSum += i
					break
				}
			}

		}
	}

	return totalSum, nil
}

func hasSequence(val string, positions int) bool {
	first := val[:positions]

	for i := positions; i < len(val); i += positions {
		if i+positions > len(val) {
			return false
		}

		next := val[i : i+positions]

		if first != next {
			return false
		}
	}

	return true
}
