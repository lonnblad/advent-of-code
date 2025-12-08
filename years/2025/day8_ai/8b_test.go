package day8_ai_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day8_ai"
	"github.com/stretchr/testify/require"
)

//go:embed 8b_input.example.txt
var exampleInput8B string

const expectedExampleOutput8B = 25272

func Test_8B_Example(t *testing.T) {
	actualOutput8B, actualErr := day8_ai.Day8B(exampleInput8B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput8B, actualOutput8B)
}

//go:embed 8b_input.txt
var input8B string

const expectedOutput8B = 169521198

func Test_8B_Input(t *testing.T) {
	actualOutput8B, actualErr := day8_ai.Day8B(input8B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput8B, actualOutput8B)
}

