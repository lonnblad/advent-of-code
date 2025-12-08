package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func Day2A(input string) (_ int, err error) {
	intervals := strings.Split(input, ",")

	var totalSum int

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, interval := range intervals {
		matches := re.FindStringSubmatch(interval)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])

		for i := start; i <= end; i++ {
			val := strconv.Itoa(i)

			firstHalf := val[:len(val)/2]
			secondHalf := val[len(val)/2:]

			if firstHalf == secondHalf {
				totalSum += i
			}
		}
	}

	return totalSum, nil
}
