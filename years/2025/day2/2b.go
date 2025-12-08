package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
)

func Day2B(input string) (_ int, err error) {
	intervals := strings.Split(input, ",")
	var totalSum int

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, intervalStr := range intervals {
		matches := re.FindStringSubmatch(intervalStr)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])

		iv := interval.Interval{Start: start, End: end}

		for i := iv.Start; i <= iv.End; i++ {
			if hasRepeatingSequence(i) {
				totalSum += i
			}
		}
	}

	return totalSum, nil
}

func hasRepeatingSequence(n int) bool {
	val := strconv.Itoa(n)
	maxPatternLen := len(val) / 2

	for patternLen := 1; patternLen <= maxPatternLen; patternLen++ {
		if isRepeatingPattern(val, patternLen) {
			return true
		}
	}

	return false
}

func isRepeatingPattern(val string, patternLen int) bool {
	if len(val)%patternLen != 0 {
		return false
	}

	pattern := val[:patternLen]

	for i := patternLen; i < len(val); i += patternLen {
		segment := val[i : i+patternLen]
		if segment != pattern {
			return false
		}
	}

	return true
}
