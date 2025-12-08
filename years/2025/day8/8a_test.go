package day8_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day8"
	"github.com/stretchr/testify/require"
)

//go:embed 8a_input.example.txt
var exampleInput1A string

const expectedExampleOutput1A = 40

func Test_1A_Example(t *testing.T) {
	actualOutput1A, actualErr := day8.Day8A(exampleInput1A, 10)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput1A, actualOutput1A)
}

//go:embed 8a_input.txt
var input1A string

const expectedOutput1A = 164475

func Test_1A_Input(t *testing.T) {
	actualOutput1A, actualErr := day8.Day8A(input1A, 1000)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput1A, actualOutput1A)
}
