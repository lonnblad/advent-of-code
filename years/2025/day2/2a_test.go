package day2_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day2"
	"github.com/stretchr/testify/require"
)

//go:embed 2a_input.example.txt
var exampleInput1A string

const expectedExampleOutput1A = 1227775554

func Test_1A_Example(t *testing.T) {
	actualOutput1A, actualErr := day2.Day2A(exampleInput1A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput1A, actualOutput1A)
}

//go:embed 2a_input.txt
var input1A string

const expectedOutput1A = 26255179562

func Test_1A_Input(t *testing.T) {
	actualOutput1A, actualErr := day2.Day2A(input1A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput1A, actualOutput1A)
}
