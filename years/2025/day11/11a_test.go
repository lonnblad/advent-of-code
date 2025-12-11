package day11_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day11"
	"github.com/stretchr/testify/require"
)

//go:embed 11a_input.example.txt
var exampleInput11A string

const expectedExampleOutput11A = 5

func Test_11A_Example(t *testing.T) {
	actualOutput11A, actualErr := day11.Day11A(exampleInput11A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput11A, actualOutput11A)
}

//go:embed 11a_input.txt
var input11A string

const expectedOutput11A = 696

func Test_11A_Input(t *testing.T) {
	actualOutput11A, actualErr := day11.Day11A(input11A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput11A, actualOutput11A)
}
