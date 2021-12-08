package aoc2021day3

import (
	_ "embed"
	"fmt"
	//"github.com/vitorf7/advent-of-code/internal/puzzle_helpers"

	"github.com/vitorf7/advent-of-code/internal/puzzle_helpers"
)

//go:embed input.txt
var day3PuzzleInput string

type Day3 struct {
	input string
}

func New() *Day3 {
	return &Day3{
		input: day3PuzzleInput,
	}
}

func (d *Day3) Solve() (string, error) {
	gamma, epsilon, powerConsumption, err := d.puzzle1(puzzle_helpers.SplitStringToSlice(d.input, "\n"))
	if err != nil {
		return "Error Processing Puzzle 1", err
	}

	return fmt.Sprintf(
		`Day 3 Puzzle 1 Result: Gamma = %d, Epsilon = %d, Power Consumption = %d
Day 3 Puzzle 2 Result: %d`,
		gamma, epsilon, powerConsumption,
		0), nil
}

func (d *Day3) puzzle1(puzzleInput []string) (int64, int64, int64, error) {
	bitSlice := make([]int, len(puzzleInput[0]))
	for _, binaryLine := range puzzleInput {
		for bitIndex, bit := range binaryLine {
			if bit == '1' {
				// Add to the slice
				bitSlice[bitIndex]++
			}
		}
	}

	gamma := make([]int, len(bitSlice))
	epsilon := make([]int, len(bitSlice))
	for index, bitCount := range bitSlice {
		if bitCount > (len(puzzleInput) / 2) {
			gamma[index] = 1
			epsilon[index] = 0
		} else {
			gamma[index] = 0
			epsilon[index] = 1
		}
	}

	gammaDecimal := d.binaryToDecimal(gamma)
	epsilonDecimal := d.binaryToDecimal(epsilon)

	return gammaDecimal, epsilonDecimal, gammaDecimal * epsilonDecimal, nil
}

func (d *Day3) puzzle2(puzzleInput []string) (int64, int64, int64, error) {
	return 0, 0, 0, nil
}

func (d *Day3) binaryToDecimal(binNumbers []int) int64 {
	// Assuming that number binNumbers 0,1s
	// Used to store result
	var result int64 = 0
	var bit int = 0

	// Execute given binNumbers in reverse order
	for i := len(binNumbers) - 1; i >= 0; i-- {
		if binNumbers[i] == 1 {
			// When get binary 1
			result += (1 << (bit))
		}
		// Count number of bits
		bit++
	}

	return result
}
