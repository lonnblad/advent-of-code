package day8_ai

import (
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code/util/geometry"
	"github.com/lonnblad/advent-of-code/util/strconv"
	"github.com/lonnblad/advent-of-code/util/unionfind"
)

func Day8A(input string, noOfConnections int) (_ int, err error) {
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

	for _, edge := range edges[:noOfConnections] {
		uf.Union(edge.From, edge.To)
	}

	sizes := uf.ComponentSizes()

	return sizes[0] * sizes[1] * sizes[2], nil
}

type Edge struct {
	From   geometry.Point3D
	To     geometry.Point3D
	Weight float64
}
