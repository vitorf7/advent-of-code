package aoc2021day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorf7/advent-of-code/internal/puzzle_helpers"
)

func TestDay1_puzzle1(t *testing.T) {
	puzzleInput := `199
200
208
210
200
207
240
269
260
263`

	expected := 7

	inputIntSlice, err := puzzle_helpers.ConvertMultiLineStringToIntSlice(puzzleInput)
	if err != nil {
		t.Error(err)
	}

	puzzle := New()
	puzzle.input = puzzleInput
	actual, err := puzzle.puzzle1(inputIntSlice)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestDay1_puzzle2(t *testing.T) {
	puzzleInput := `199
200
208
210
200
207
240
269
260
263`

	expected := 5

	inputIntSlice, err := puzzle_helpers.ConvertMultiLineStringToIntSlice(puzzleInput)
	if err != nil {
		t.Error(err)
	}

	puzzle := New()
	puzzle.input = puzzleInput
	actual, err := puzzle.puzzle2(inputIntSlice, 3)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestDay1_Solve(t *testing.T) {
	puzzleInput := `199
200
208
210
200
207
240
269
260
263`
	expected := `Day 1 Puzzle 1 Result: Number of increases 7
Day 1 Puzzle 2 Result: Number of increases 5`

	puzzle := New()
	puzzle.input = puzzleInput

	actual, err := puzzle.Solve()

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}
