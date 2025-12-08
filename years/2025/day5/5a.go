package day5

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
	"github.com/lonnblad/advent-of-code/util/strconv"
)

func Day5A(input string) (_ int, err error) {
	sections := strings.Split(input, "\n\n")
	freshRanges := parseFreshRanges(sections[0])
	usedIngredients := parseUsedIngredients(sections[1])

	var totalSum int

	for _, ingredient := range usedIngredients {
		if isInAnyRange(ingredient, freshRanges) {
			totalSum++
		}
	}

	return totalSum, nil
}

func parseFreshRanges(input string) []interval.Interval {
	lines := strings.Split(input, "\n")
	ranges := make([]interval.Interval, len(lines))

	for idx, line := range lines {
		ranges[idx] = interval.ParseInterval(line)
	}

	return ranges
}

func parseUsedIngredients(input string) []int {
	lines := strings.Split(input, "\n")
	ingredients := make([]int, len(lines))

	for idx, line := range lines {
		ingredients[idx] = strconv.MustParseInt(line)
	}

	return ingredients
}

func isInAnyRange(value int, ranges []interval.Interval) bool {
	for _, r := range ranges {
		if r.Contains(value) {
			return true
		}
	}
	return false
}
