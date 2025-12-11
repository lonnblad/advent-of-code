package graph

import (
	"strings"
)

// Graph represents a directed graph with string nodes
type Graph struct {
	adjacency map[string][]string
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		adjacency: make(map[string][]string),
	}
}

// AddEdge adds a directed edge from 'from' to 'to'.
// Both nodes are added to the graph if they don't already exist.
func (g *Graph) AddEdge(from, to string) {
	g.adjacency[from] = append(g.adjacency[from], to)

	// Ensure 'to' node exists in the graph (even if it has no outgoing edges)
	if _, exists := g.adjacency[to]; !exists {
		g.adjacency[to] = nil
	}
}

// AddEdges adds multiple edges from 'from' to all nodes in 'tos'
func (g *Graph) AddEdges(from string, tos []string) {
	for _, to := range tos {
		g.AddEdge(from, to)
	}
}

// GetNeighbors returns all neighbors of a node.
// Returns an empty slice if the node doesn't exist or has no neighbors.
func (g *Graph) GetNeighbors(node string) []string {
	neighbors := g.adjacency[node]
	if neighbors == nil {
		return []string{}
	}

	return neighbors
}

// HasNode checks if a node exists in the graph
func (g *Graph) HasNode(node string) bool {
	_, exists := g.adjacency[node]
	return exists
}

// CountPaths counts paths from 'from' to 'to' in a DAG.
// If required nodes are specified, only counts paths that pass through all of them in any order.
// Uses DFS with memoization. Assumes the graph is acyclic.
// Returns 0 if no path exists, or the number of distinct paths.
// If no required nodes are specified, counts all paths from 'from' to 'to'.
func (g *Graph) CountPaths(from, to string, required ...string) int {
	if !g.HasNode(from) || !g.HasNode(to) {
		return 0
	}

	// Convert required to a set for quick lookup
	requiredSet := make(map[string]bool, len(required))
	for _, node := range required {
		requiredSet[node] = true

		// Ensure all required nodes exist
		if !g.HasNode(node) {
			return 0
		}
	}

	// Optimize: if no required nodes, use simpler DFS without tracking visited required nodes
	if len(requiredSet) == 0 {
		memo := make(map[string]int)

		var dfs func(node string) int
		dfs = func(node string) int {
			if node == to {
				return 1
			}

			// Check memoization
			if count, exists := memo[node]; exists {
				return count
			}

			totalPaths := 0
			for _, neighbor := range g.GetNeighbors(node) {
				totalPaths += dfs(neighbor)
			}

			memo[node] = totalPaths
			return totalPaths
		}

		return dfs(from)
	}

	// Memoization: key is "node:visitedRequired" where visitedRequired is a string
	// representation of visited required nodes
	memo := make(map[string]int)

	// Helper to create a key for memoization
	createKey := func(node string, visitedRequired map[string]bool) string {
		key := node + ":"

		// Create consistent key by iterating required nodes in order
		for _, req := range required {
			if visitedRequired[req] {
				key += req + ","
			}
		}

		return key
	}

	var dfs func(node string, visitedRequired map[string]bool) int
	dfs = func(node string, visitedRequired map[string]bool) int {
		// Check if we've reached the destination
		if node == to {
			// Verify all required nodes were visited
			for req := range requiredSet {
				if !visitedRequired[req] {
					return 0
				}
			}

			return 1
		}

		// Check memoization
		key := createKey(node, visitedRequired)
		if count, exists := memo[key]; exists {
			return count
		}

		// Mark current node as visited if it's required
		newVisitedRequired := make(map[string]bool, len(visitedRequired)+1)
		for k, v := range visitedRequired {
			newVisitedRequired[k] = v
		}

		if requiredSet[node] {
			newVisitedRequired[node] = true
		}

		totalPaths := 0
		for _, neighbor := range g.GetNeighbors(node) {
			totalPaths += dfs(neighbor, newVisitedRequired)
		}

		memo[key] = totalPaths
		return totalPaths
	}

	initialVisited := make(map[string]bool)
	return dfs(from, initialVisited)
}

// ParseGraph parses a graph from input lines in the format:
// "node: neighbor1 neighbor2 ..."
// Each line represents edges from 'node' to its neighbors
func ParseGraph(lines []string) *Graph {
	g := NewGraph()

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Split on ":"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		from := strings.TrimSpace(parts[0])
		rest := strings.TrimSpace(parts[1])

		// Parse neighbors (split on whitespace)
		neighbors := strings.Fields(rest)

		g.AddEdges(from, neighbors)
	}

	return g
}
