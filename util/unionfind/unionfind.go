package unionfind

import "sort"

// UnionFind is a generic disjoint set union data structure that efficiently
// tracks which elements belong to the same set/component.
//
// It supports two main operations:
//   - Find: Determine which set an element belongs to (O(α(n)) amortized)
//   - Union: Merge two sets together (O(α(n)) amortized)
//
// T must be a comparable type (can be used as a map key).
type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	size   map[T]int
}

// New creates a new UnionFind data structure initialized with the given elements.
// Each element starts in its own set.
func New[T comparable](elements []T) *UnionFind[T] {
	uf := &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		size:   make(map[T]int),
	}

	for _, elem := range elements {
		uf.parent[elem] = elem
		uf.rank[elem] = 0
		uf.size[elem] = 1
	}

	return uf
}

// Find returns the root/representative of the set containing x.
// Uses path compression to optimize future Find operations.
func (uf *UnionFind[T]) Find(x T) T {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}

	return uf.parent[x]
}

// Union merges the sets containing x and y.
// Returns true if the sets were merged, false if they were already in the same set.
// Uses union by rank to keep trees balanced.
func (uf *UnionFind[T]) Union(x, y T) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false // Already in same set
	}

	// Union by rank: attach smaller tree to larger tree
	if uf.rank[rootX] < uf.rank[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]

	if uf.rank[rootX] == uf.rank[rootY] {
		uf.rank[rootX]++
	}

	return true
}

// Size returns the size of the set containing x.
func (uf *UnionFind[T]) Size(x T) int {
	root := uf.Find(x)
	return uf.size[root]
}

// ComponentSizes returns the sizes of all components, sorted in descending order.
func (uf *UnionFind[T]) ComponentSizes() []int {
	sizes := make(map[T]int)

	for node := range uf.parent {
		root := uf.Find(node)
		sizes[root] = uf.size[root]
	}

	result := make([]int, 0, len(sizes))

	for _, size := range sizes {
		result = append(result, size)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result
}

// Connected returns true if x and y are in the same set.
func (uf *UnionFind[T]) Connected(x, y T) bool {
	return uf.Find(x) == uf.Find(y)
}

// ComponentCount returns the number of distinct components.
func (uf *UnionFind[T]) ComponentCount() int {
	roots := make(map[T]bool)

	for node := range uf.parent {
		root := uf.Find(node)
		roots[root] = true
	}

	return len(roots)
}

