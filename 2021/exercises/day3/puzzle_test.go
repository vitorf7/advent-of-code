package aoc2021day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorf7/advent-of-code/2021/internal/puzzle_helpers"
)

func TestDay3_puzzle1(t *testing.T) {
	testInput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	expectedGamma := int64(22)
	expectedEpsilon := int64(9)
	expectedPowerConsumption := int64(198)

	puzzle := New()
	puzzle.input = testInput
	actualGamma, actualEpsilon, actualPowerConsumption, actualError := puzzle.puzzle1(puzzle_helpers.SplitStringToSlice(testInput, "\n"))

	assert.Equal(t, expectedGamma, actualGamma)
	assert.Equal(t, expectedEpsilon, actualEpsilon)
	assert.Equal(t, expectedPowerConsumption, actualPowerConsumption)
	assert.NoError(t, actualError)
}

func TestDay3_puzzle2(t *testing.T) {
	testInput := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	expectedOxygenRating := int64(23)
	expectedCO2Rating := int64(10)
	expectedLifeSupportRating := int64(230)

	puzzle := New()
	puzzle.input = testInput
	actualOxygenRating, actualCO2Rating, actualPowerLifeSupportRating, actualError := puzzle.puzzle2(
		puzzle_helpers.SplitStringToSlice(testInput, "\n"),
	)

	assert.Equal(t, expectedOxygenRating, actualOxygenRating)
	assert.Equal(t, expectedCO2Rating, actualCO2Rating)
	assert.Equal(t, expectedLifeSupportRating, actualPowerLifeSupportRating)
	assert.NoError(t, actualError)
}
