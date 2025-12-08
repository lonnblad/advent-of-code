package day5

import (
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code/util/interval"
)

func Day5B(input string) (_ int, err error) {
	sections := strings.Split(input, "\n\n")
	freshRanges := parseFreshRanges(sections[0])
	mergedRanges := mergeIntervals(freshRanges)

	var totalSum int

	for _, r := range mergedRanges {
		totalSum += r.Size()
	}

	return totalSum, nil
}

// mergeIntervals merges all overlapping intervals using a single-pass algorithm.
// First sorts intervals by start, then merges overlapping ones in one pass.
func mergeIntervals(intervals []interval.Interval) []interval.Interval {
	if len(intervals) == 0 {
		return intervals
	}

	// Sort intervals by start position
	sorted := make([]interval.Interval, len(intervals))
	copy(sorted, intervals)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Start < sorted[j].Start
	})

	// Merge overlapping intervals in a single pass
	merged := []interval.Interval{sorted[0]}

	for i := 1; i < len(sorted); i++ {
		last := &merged[len(merged)-1]
		current := sorted[i]

		if last.HasOverlap(current) {
			// Merge with the last interval
			*last = last.Union(current)
		} else {
			// Add as a new interval
			merged = append(merged, current)
		}
	}

	return merged
}
