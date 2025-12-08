package matrix

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/strconv"
)

func IntMatrix(a string) [][]int {
	rows := strings.Split(a, "\n")
	result := make([][]int, len(rows))

	for i, row := range rows {
		cells := strings.Split(row, "")
		result[i] = make([]int, len(cells))

		for j, cell := range cells {
			result[i][j] = strconv.MustParseInt(cell)
		}
	}

	return result
}

func StringMatrix(a string) [][]string {
	rows := strings.Split(a, "\n")
	result := make([][]string, len(rows))

	for i, row := range rows {
		result[i] = strings.Split(row, "")
	}

	return result
}

type Matrix struct {
	Values  [][]string
	noRows  int
	noCells int
}

func NewMatrix(a string) Matrix {
	rows := strings.Split(a, "\n")
	result := make([][]string, len(rows))

	for i, row := range rows {
		result[i] = strings.Split(row, "")
	}

	return Matrix{
		Values:  result,
		noRows:  len(result),
		noCells: len(result[0]),
	}
}

// IsWithinBounds checks if the given coordinate is within the bounds of the matrix
// The coordinate is a tuple of the row and column indices
func (m Matrix) IsWithinBounds(coord [2]int) bool {
	return coord[0] >= 0 && coord[0] < m.noRows && coord[1] >= 0 && coord[1] < m.noCells
}

// GetAdjacentCells returns the adjacent cells of the given coordinate
// The coordinate is a tuple of the row and column indices
// The adjacent cells are the cells that are adjacent to the given coordinate
// The adjacent cells are returned in the order of top right, top left, bottom right, bottom left, top, bottom, left, right
func (m Matrix) GetAdjacentCells(coord [2]int) []string {
	adjacentCells := make([]string, 0, 8)

	// Top right
	if m.IsWithinBounds([2]int{coord[0] + 1, coord[1] + 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]+1][coord[1]+1])
	}

	// Top left
	if m.IsWithinBounds([2]int{coord[0] + 1, coord[1] - 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]+1][coord[1]-1])
	}

	// Bottom right
	if m.IsWithinBounds([2]int{coord[0] - 1, coord[1] + 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]-1][coord[1]+1])
	}

	// Bottom left
	if m.IsWithinBounds([2]int{coord[0] - 1, coord[1] - 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]-1][coord[1]-1])
	}

	// Top
	if m.IsWithinBounds([2]int{coord[0], coord[1] + 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]][coord[1]+1])
	}

	// Bottom
	if m.IsWithinBounds([2]int{coord[0], coord[1] - 1}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]][coord[1]-1])
	}

	// Left
	if m.IsWithinBounds([2]int{coord[0] - 1, coord[1]}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]-1][coord[1]])
	}

	// Right
	if m.IsWithinBounds([2]int{coord[0] + 1, coord[1]}) {
		adjacentCells = append(adjacentCells, m.Values[coord[0]+1][coord[1]])
	}

	return adjacentCells
}

func (m Matrix) Find(target string) (int, int) {
	for i, row := range m.Values {
		for j, cell := range row {
			if cell == target {
				return i, j
			}
		}
	}

	return -1, -1
}

func (m Matrix) String() string {
	var result strings.Builder

	for _, row := range m.Values {
		for _, cell := range row {
			result.WriteString(cell)
		}
		result.WriteString("\n")
	}

	return result.String()
}

// Transpose transposes a slice of strings (rows become columns)
func Transpose(a []string) []string {
	newArr := make([][]byte, len(a[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	newStringArr := make([]string, len(newArr))
	for idx, row := range newArr {
		newStringArr[idx] = string(row)
	}

	return newStringArr
}

// Rotate45Degrees rotates a matrix 45 degrees
func Rotate45Degrees(a []string) []string {
	n := len(a)
	if n == 0 {
		return nil
	}

	m := len(a[0])
	if m == 0 {
		return nil
	}

	// The size of the resulting array will be n + m - 1
	result := make([]string, n+m-1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i+j] += string(a[i][j])
		}
	}

	return result
}

