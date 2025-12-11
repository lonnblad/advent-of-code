package day11

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/graph"
)

func Day11A(input string) (_ int, err error) {
	lines := strings.Split(input, "\n")

	graph := graph.ParseGraph(lines)
	paths := graph.CountPaths("you", "out")

	return paths, nil
}
