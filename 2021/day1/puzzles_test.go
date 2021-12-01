package aoc2021day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
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

	actual, err := Puzzle1(puzzleInput)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestPuzzle2(t *testing.T) {
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

	actual, err := Puzzle2(puzzleInput, 3)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}
