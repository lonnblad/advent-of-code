package day2_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day2"
	"github.com/stretchr/testify/require"
)

//go:embed 2b_input.example.txt
var exampleInput1B string

const expectedExampleOutput1B = 4174379265

func Test_1B_Example(t *testing.T) {
	actualOutput1B, actualErr := day2.Day2B(exampleInput1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput1B, actualOutput1B)
}

//go:embed 2b_input.txt
var input1B string

const expectedOutput1B = 31680313976

func Test_1B_Input(t *testing.T) {
	actualOutput1B, actualErr := day2.Day2B(input1B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput1B, actualOutput1B)
}
