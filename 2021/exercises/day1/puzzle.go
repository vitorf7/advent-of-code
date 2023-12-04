package aoc2021day1

import (
	_ "embed"
	"fmt"

	"github.com/vitorf7/advent-of-code/2021/internal/puzzle_helpers"
)

//go:embed input.txt
var day1PuzzleInput string

type Day1 struct {
	input string
}

func New() *Day1 {
	return &Day1{
		input: day1PuzzleInput,
	}
}

func (d *Day1) Solve() (string, error) {
	inputIntSlice, err := puzzle_helpers.ConvertMultiLineStringToIntSlice(d.input)
	if err != nil {
		return "", err
	}

	puzzle1Result, err := d.puzzle1(inputIntSlice)
	if err != nil {
		return "", err
	}

	puzzle2Result, err := d.puzzle2(inputIntSlice, 3)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`Day 1 Puzzle 1 Result: Number of increases %d
Day 1 Puzzle 2 Result: Number of increases %d`,
		puzzle1Result,
		puzzle2Result), nil
}

func (d *Day1) puzzle1(puzzleInput []int) (int, error) {
	increases := 0
	previousNumber := 0

	for _, currentNumber := range puzzleInput {
		if previousNumber == 0 {
			previousNumber = currentNumber
			continue
		}

		if currentNumber > previousNumber {
			increases++
		}
		previousNumber = currentNumber
	}

	return increases, nil
}

func (d *Day1) puzzle2(puzzleInput []int, windowSize int) (int, error) {
	increases := 0
	previousNumber := 0

	start := 0
	currentSum := 0
	for index, value := range puzzleInput {
		currentSum += value

		if previousNumber != 0 && currentSum > previousNumber {
			increases++
		}

		if index-start+1 == windowSize {
			previousNumber = currentSum
			currentSum -= puzzleInput[start]
			start += 1
		}
	}

	return increases, nil
}
