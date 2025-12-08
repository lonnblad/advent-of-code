package day8

import (
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code/util/geometry"
	"github.com/lonnblad/advent-of-code/util/strconv"
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

	graph := Graph{
		Nodes:          nodes,
		Edges:          make([]Edge, 0, 10),
		connectedNodes: make(map[geometry.Point3D]map[geometry.Point3D]bool),
	}

	for _, edge := range edges[:noOfConnections] {
		graph.ConnectNodes(edge.From, edge.To)
	}

	circuits := graph.FindCircuits()

	return circuits[0] * circuits[1] * circuits[2], nil
}

type Graph struct {
	Nodes []geometry.Point3D
	Edges []Edge

	connectedNodes map[geometry.Point3D]map[geometry.Point3D]bool
}

type Edge struct {
	From   geometry.Point3D
	To     geometry.Point3D
	Weight float64
}

type Circuit struct {
	Nodes map[geometry.Point3D]bool
}

func (c *Circuit) HasNodes(p1, p2 geometry.Point3D) bool {
	return c.Nodes[p1] && c.Nodes[p2]
}

func (c *Circuit) AddNode(node geometry.Point3D) {
	c.Nodes[node] = true
}

func (c *Circuit) Merge(other Circuit) {
	for node := range other.Nodes {
		c.Nodes[node] = true
	}
}

func (g *Graph) Traverse(p1, p2 geometry.Point3D) bool {
	return false
}

func (g *Graph) FindCircuits() []int {
	visitedNodes := map[geometry.Point3D]bool{}
	circuitSizes := []int{}

	for node, connectedNodes := range g.connectedNodes {
		if visitedNodes[node] {
			continue
		}

		visitedNodes[node] = true
		for connectedNode := range connectedNodes {
			visitedNodes[connectedNode] = true
		}

		circuitSizes = append(circuitSizes, len(connectedNodes))
	}

	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})

	return circuitSizes[:3]
}

func (g *Graph) ConnectNodes(p1, p2 geometry.Point3D) bool {
	if g.connectedNodes[p1][p2] || g.connectedNodes[p2][p1] {
		return false
	}

	if g.connectedNodes[p1] == nil {
		g.connectedNodes[p1] = make(map[geometry.Point3D]bool)
	}

	if g.connectedNodes[p2] == nil {
		g.connectedNodes[p2] = make(map[geometry.Point3D]bool)
	}

	g.connectedNodes[p1][p2] = true
	g.connectedNodes[p2][p1] = true

	for node := range g.connectedNodes[p1] {
		g.ConnectNodes(node, p2)
	}

	for node := range g.connectedNodes[p2] {
		g.ConnectNodes(node, p1)
	}

	return true
}
