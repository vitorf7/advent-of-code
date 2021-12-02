package aoc2021day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorf7/advent-of-code/2021/puzzle_helpers"
)

func TestDay2_puzzle1(t *testing.T) {
	puzzleInput := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	expectedHorizontal := 15
	expectedDepth := 10
	expectedTotal := 150

	inputSlice := puzzle_helpers.SplitStringToSlice(puzzleInput, "\n")

	puzzle := New()
	puzzle.input = puzzleInput
	actualHorizontal, actualDepth, actualTotal, err := puzzle.puzzle1(inputSlice)

	assert.Equal(t, expectedHorizontal, actualHorizontal)
	assert.Equal(t, expectedDepth, actualDepth)
	assert.Equal(t, expectedTotal, actualTotal)
	assert.NoError(t, err)
}

func TestDay2_puzzle2(t *testing.T) {
	puzzleInput := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	expectedHorizontal := 15
	expectedDepth := 60
	expectedTotal := 900

	inputSlice := puzzle_helpers.SplitStringToSlice(puzzleInput, "\n")

	puzzle := New()
	puzzle.input = puzzleInput
	actualHorizontal, actualDepth, actualTotal, err := puzzle.puzzle2(inputSlice)

	assert.Equal(t, expectedHorizontal, actualHorizontal)
	assert.Equal(t, expectedDepth, actualDepth)
	assert.Equal(t, expectedTotal, actualTotal)
	assert.NoError(t, err)
}
