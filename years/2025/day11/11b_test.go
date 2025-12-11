package day11_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day11"
	"github.com/stretchr/testify/require"
)

//go:embed 11b_input.example.txt
var exampleInput11B string

const expectedExampleOutput11B = 2

func Test_11B_Example(t *testing.T) {
	actualOutput11B, actualErr := day11.Day11B(exampleInput11B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput11B, actualOutput11B)
}

//go:embed 11b_input.txt
var input11B string

const expectedOutput11B = 473741288064360

func Test_11B_Input(t *testing.T) {
	actualOutput11B, actualErr := day11.Day11B(input11B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput11B, actualOutput11B)
}
