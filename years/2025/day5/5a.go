package day5

import (
	"strconv"
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
)

func Day5A(input string) (_ int, err error) {
	var totalSum int

	sections := strings.Split(input, "\n\n")

	fresh := strings.Split(sections[0], "\n")
	freshRanges := make([]interval.Interval, len(fresh))

	for idx, r := range fresh {
		freshRanges[idx] = interval.ParseInterval(r)
	}

	used := strings.Split(sections[1], "\n")
	usedIngredients := make([]int, len(used))

	for idx, ingredient := range used {
		usedIngredients[idx], _ = strconv.Atoi(ingredient)

		for _, freshRange := range freshRanges {
			if freshRange.Contains(usedIngredients[idx]) {
				totalSum++
				break
			}
		}
	}

	return totalSum, nil
}
