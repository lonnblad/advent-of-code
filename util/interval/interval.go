package interval

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/strconv"
)

type Interval struct {
	Start int
	End   int
}

func ParseInterval(input string) Interval {
	parts := strings.Split(input, "-")

	return Interval{
		Start: strconv.MustParseInt(parts[0]),
		End:   strconv.MustParseInt(parts[1]),
	}
}

func (i Interval) Contains(value int) bool {
	return value >= i.Start && value <= i.End
}

func (i Interval) AllValues() []int {
	values := make([]int, i.End-i.Start+1)

	for idx := i.Start; idx <= i.End; idx++ {
		values[idx-i.Start] = idx
	}

	return values
}

func (i Interval) Size() int {
	return i.End - i.Start + 1
}

func (i Interval) HasOverlap(other Interval) bool {
	return i.Start <= other.End && i.End >= other.Start
}

func (i Interval) Union(other Interval) Interval {
	return Interval{
		Start: min(i.Start, other.Start),
		End:   max(i.End, other.End),
	}
}

func (i Interval) Intersection(other Interval) Interval {
	return Interval{
		Start: max(i.Start, other.Start),
		End:   min(i.End, other.End),
	}
}

