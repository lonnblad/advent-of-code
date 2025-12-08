package day8_ai

import (
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code/util/geometry"
	"github.com/lonnblad/advent-of-code/util/strconv"
	"github.com/lonnblad/advent-of-code/util/unionfind"
)

func Day8B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	nodes := make([]geometry.Point3D, len(rows))
	edges := make([]Edge, 0)

	for idx, row := range rows {
		parts := strings.Split(row, ",")
		x := strconv.MustParseInt(parts[0])
		y := strconv.MustParseInt(parts[1])
		z := strconv.MustParseInt(parts[2])

		nodes[idx] = geometry.Point3D{X: x, Y: y, Z: z}

		for jdx := 0; jdx < idx; jdx++ {
			edge := Edge{
				From:   nodes[idx],
				To:     nodes[jdx],
				Weight: nodes[idx].Distance(nodes[jdx]),
			}

			edges = append(edges, edge)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	uf := unionfind.New(nodes)
	var lastEdge Edge

	for _, edge := range edges {
		if uf.Union(edge.From, edge.To) {
			lastEdge = edge
			// Check if fully connected (all nodes in one component)
			if uf.Size(nodes[0]) == len(nodes) {
				break
			}
		}
	}

	return lastEdge.From.X * lastEdge.To.X, nil
}
