package day11

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/graph"
)

func Day11B(input string) (_ int, err error) {
	lines := strings.Split(input, "\n")

	graph := graph.ParseGraph(lines)
	paths := graph.CountPaths("svr", "out", "dac", "fft")

	return paths, nil
}
