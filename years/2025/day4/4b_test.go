package day4_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day4"
	"github.com/stretchr/testify/require"
)

//go:embed 4b_input.example.txt
var exampleInput1B string

const expectedExampleOutput1B = 43

func Test_1B_Example(t *testing.T) {
	actualOutput1B, actualErr := day4.Day4B(exampleInput1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput1B, actualOutput1B)
}

//go:embed 4b_input.txt
var input1B string

const expectedOutput1B = 8277

func Test_1B_Input(t *testing.T) {
	actualOutput1B, actualErr := day4.Day4B(input1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput1B, actualOutput1B)
}
