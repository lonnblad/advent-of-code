package day6_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day6"
	"github.com/stretchr/testify/require"
)

//go:embed 6b_input.example.txt
var exampleInput1B string

const expectedExampleOutput1B = 3263827

func Test_1B_Example(t *testing.T) {
	actualOutput1B, actualErr := day6.Day6B(exampleInput1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput1B, actualOutput1B)
}

//go:embed 6b_input.txt
var input1B string

const expectedOutput1B = 10389131401929

func Test_1B_Input(t *testing.T) {
	actualOutput1B, actualErr := day6.Day6B(input1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput1B, actualOutput1B)
}
