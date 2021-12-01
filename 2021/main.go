package main

import (
	_ "embed"
	"fmt"

	aoc2021day1 "github.com/vitorf7/advent-of-code/2021/day1"
)

//go:embed day1/input.txt
var day1PuzzleInput string

func main() {
	day1Puzzle1Result, err := aoc2021day1.Puzzle1(day1PuzzleInput)
	if err != nil {
		fmt.Printf("There was an error with Day 1 Puzzle 1: %+v", err)
		return
	}
	fmt.Printf("Day 1 Puzzle 1 Result: Number of increases %d\n", day1Puzzle1Result)
	day1Puzzle2Result, err := aoc2021day1.Puzzle2(day1PuzzleInput, 3)
	if err != nil {
		fmt.Printf("There was an error with Day 1 Puzzle 2: %+v", err)
		return
	}
	fmt.Printf("Day 1 Puzzle 2 Result: Number of increases %d\n", day1Puzzle2Result)
}
