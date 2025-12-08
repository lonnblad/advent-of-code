package day8

import (
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code/util/geometry"
	"github.com/lonnblad/advent-of-code/util/strconv"
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

	// fmt.Println(edges[:10])

	graph := Graph{
		Nodes:          nodes,
		Edges:          make([]Edge, 0, 10),
		connectedNodes: make(map[geometry.Point3D]map[geometry.Point3D]bool),
	}

	var lastConnectedEdge Edge

	for _, edge := range edges {
		if graph.ConnectNodes(edge.From, edge.To) {
			lastConnectedEdge = edge
		}
	}

	// fmt.Println(lastConnectedEdge)

	// fmt.Println("--------------------------------")
	// fmt.Println("Connected Nodes")
	// fmt.Println("--------------------------------")
	// for node, connectedNodes := range graph.connectedNodes {
	// 	fmt.Println(node, len(connectedNodes))
	// 	fmt.Println(connectedNodes)
	// 	fmt.Println("--------------------------------")
	// }

	// circuits := graph.FindCircuits()
	// fmt.Println(circuits)

	return lastConnectedEdge.From.X * lastConnectedEdge.To.X, nil
}
