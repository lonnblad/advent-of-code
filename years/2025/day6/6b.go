package day6

import (
	"strings"

	"github.com/lonnblad/advent-of-code/util/strconv"
)

func Day6B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")
	lastRow := rows[len(rows)-1]
	dataRows := rows[:len(rows)-1]

	operators := propagateOperators(lastRow)

	var totalSum int
	var currentSum int

	for colIdx := range lastRow {
		if operators[colIdx] == "" {
			// Empty operator marks end of a group, add accumulated sum
			totalSum += currentSum
			currentSum = 0
			continue
		}

		number := extractColumnNumber(dataRows, colIdx)
		currentSum = applyOperator(currentSum, number, operators[colIdx])
	}

	// Add the last accumulated sum
	totalSum += currentSum

	return totalSum, nil
}

// propagateOperators propagates operators from the first row through columns.
// When an operator is found, it replaces the previous operator in that position.
// Operators propagate forward through empty cells.
func propagateOperators(firstRow string) []string {
	cells := strings.Split(firstRow, "")
	operators := make([]string, len(cells))

	for idx, cell := range cells {
		if idx == 0 {
			operators[idx] = cell
			continue
		}

		if cell == "+" || cell == "*" {
			operators[idx] = cell
			// Clear the previous operator when a new one is found
			if idx > 0 {
				operators[idx-1] = ""
			}
			continue
		}

		// Propagate the previous operator forward
		operators[idx] = operators[idx-1]
	}

	return operators
}

// extractColumnNumber builds a number by concatenating digits from a column.
func extractColumnNumber(rows []string, colIdx int) int {
	var columnDigits string

	for _, row := range rows {
		cells := strings.Split(row, "")
		if colIdx < len(cells) {
			columnDigits += cells[colIdx]
		}
	}

	return strconv.MustParseInt(strings.TrimSpace(columnDigits))
}

// applyOperator applies the given operator to the current sum and number.
// For multiplication, initializes with the number if sum is 0.
func applyOperator(currentSum, number int, operator string) int {
	switch operator {
	case "+":
		return currentSum + number
	case "*":
		if currentSum == 0 {
			return number
		}
		return currentSum * number
	default:
		return currentSum
	}
}
