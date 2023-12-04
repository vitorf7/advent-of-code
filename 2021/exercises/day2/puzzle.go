package aoc2021day2

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/vitorf7/advent-of-code/2021/internal/puzzle_helpers"
)

//go:embed input.txt
var day2PuzzleInput string

type Day2 struct {
	input string
}

func New() *Day2 {
	return &Day2{
		input: day2PuzzleInput,
	}
}

func (d *Day2) Solve() (string, error) {
	puzzle1Horizontal, puzzle1Depth, puzzle1Total, err := d.puzzle1(puzzle_helpers.SplitStringToSlice(d.input, "\n"))
	if err != nil {
		return "", err
	}

	puzzle2Horizontal, puzzle2Depth, puzzle2Total, err := d.puzzle2(puzzle_helpers.SplitStringToSlice(d.input, "\n"))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`Day 2 Puzzle 1 Result: Horizontal: %d, Depth: %d, Total: %d
Day 2 Puzzle 2 Result: Horizontal: %d, Depth: %d, Total: %d`,
		puzzle1Horizontal, puzzle1Depth, puzzle1Total,
		puzzle2Horizontal, puzzle2Depth, puzzle2Total), nil
}

func (d *Day2) puzzle1(puzzleInput []string) (int, int, int, error) {
	navigationMapSlice, err := d.getNavigationMap(puzzleInput)
	if err != nil {
		return 0, 0, 0, err
	}

	horizontalPosition := 0
	depthPosition := 0
	for _, instructionMap := range navigationMapSlice {
		increaseHorizontalValue, moveForward := instructionMap["forward"]
		if moveForward {
			horizontalPosition += increaseHorizontalValue
		}

		increaseDepthValue, moveDown := instructionMap["down"]
		if moveDown {
			depthPosition += increaseDepthValue
		}

		decreaseDepthValue, moveUp := instructionMap["up"]
		if moveUp {
			depthPosition -= decreaseDepthValue
		}
	}

	return horizontalPosition, depthPosition, horizontalPosition * depthPosition, nil
}

func (d *Day2) puzzle2(puzzleInput []string) (int, int, int, error) {
	navigationMapSlice, err := d.getNavigationMap(puzzleInput)
	if err != nil {
		return 0, 0, 0, err
	}

	horizontalPosition := 0
	depthPosition := 0
	aim := 0
	for _, instructionMap := range navigationMapSlice {
		increaseHorizontalValue, moveForward := instructionMap["forward"]
		if moveForward {
			horizontalPosition += increaseHorizontalValue
			depthPosition = depthPosition + (aim * increaseHorizontalValue)
		}

		increaseDepthValue, moveDown := instructionMap["down"]
		if moveDown {
			aim += increaseDepthValue
		}

		decreaseDepthValue, moveUp := instructionMap["up"]
		if moveUp {
			aim -= decreaseDepthValue
		}
	}

	return horizontalPosition, depthPosition, horizontalPosition * depthPosition, nil
}

func (d *Day2) getNavigationMap(input []string) ([]map[string]int, error) {
	navigationMap := make([]map[string]int, len(input))
	for i, line := range input {
		instructionMap := make(map[string]int)
		instructionSlice := puzzle_helpers.SplitStringToSlice(line, " ")
		instructionInt, err := strconv.Atoi(instructionSlice[1])
		if err != nil {
			return nil, err
		}
		instructionMap[instructionSlice[0]] = instructionInt
		navigationMap[i] = instructionMap
	}

	return navigationMap, nil
}
