package day5

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
)

func Day5B(input string) (_ int, err error) {
	var totalSum int

	sections := strings.Split(input, "\n\n")

	fresh := strings.Split(sections[0], "\n")
	freshRanges := make([]interval.Interval, len(fresh))

	for idx, r := range fresh {
		freshRanges[idx] = interval.ParseInterval(r)
	}

	latestMergedRanges := freshRanges

	for {
		foundMerge := false
		mergedRanges := make([]interval.Interval, 0, len(latestMergedRanges))
		mergedIndexes := make(map[int]bool)

		for idx := 0; idx < len(latestMergedRanges); idx++ {
			if mergedIndexes[idx] {
				continue
			}

			mergedRange := latestMergedRanges[idx]

			for jdx := idx + 1; jdx < len(latestMergedRanges); jdx++ {
				if mergedIndexes[jdx] {
					continue
				}

				if mergedRange.HasOverlap(latestMergedRanges[jdx]) {
					mergedIndexes[idx] = true
					mergedIndexes[jdx] = true
					mergedRange = mergedRange.Union(latestMergedRanges[jdx])

					foundMerge = true
				}
			}

			mergedRanges = append(mergedRanges, mergedRange)
		}

		latestMergedRanges = mergedRanges

		if !foundMerge {
			break
		}
	}

	for _, mergedRange := range latestMergedRanges {
		totalSum += mergedRange.Size()
	}

	return totalSum, nil
}
