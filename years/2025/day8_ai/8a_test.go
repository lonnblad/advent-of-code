package day8_ai_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day8_ai"
	"github.com/stretchr/testify/require"
)

//go:embed 8a_input.example.txt
var exampleInput8A string

const expectedExampleOutput8A = 40

func Test_8A_Example(t *testing.T) {
	actualOutput8A, actualErr := day8_ai.Day8A(exampleInput8A, 10)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput8A, actualOutput8A)
}

//go:embed 8a_input.txt
var input8A string

const expectedOutput8A = 164475

func Test_8A_Input(t *testing.T) {
	actualOutput8A, actualErr := day8_ai.Day8A(input8A, 1000)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput8A, actualOutput8A)
}
