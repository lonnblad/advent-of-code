package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
)

func Day2A(input string) (_ int, err error) {
	intervals := strings.Split(input, ",")
	var totalSum int

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, intervalStr := range intervals {
		matches := re.FindStringSubmatch(intervalStr)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])

		iv := interval.Interval{Start: start, End: end}

		for i := iv.Start; i <= iv.End; i++ {
			if isRepeatingNumber(i) {
				totalSum += i
			}
		}
	}

	return totalSum, nil
}

func isRepeatingNumber(n int) bool {
	val := strconv.Itoa(n)

	// Only works for even-length numbers
	if len(val)%2 != 0 {
		return false
	}

	mid := len(val) / 2
	firstHalf := val[:mid]
	secondHalf := val[mid:]

	return firstHalf == secondHalf
}
