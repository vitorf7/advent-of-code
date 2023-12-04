package exercises

import (
	"fmt"

	aoc2021day1 "github.com/vitorf7/advent-of-code/2021/exercises/day1"
	aoc2021day2 "github.com/vitorf7/advent-of-code/2021/exercises/day2"
	aoc2021day3 "github.com/vitorf7/advent-of-code/2021/exercises/day3"
)

// PuzzleSolver an interface for every puzzle
type PuzzleSolver interface {
	Solve() (string, error)
}

func PrintResults() {
	puzzles := []PuzzleSolver{
		aoc2021day1.New(),
		aoc2021day2.New(),
		aoc2021day3.New(),
	}

	fmt.Println("======= Advent of Code 2021 =======")
	for _, puzzle := range puzzles {
		result, err := puzzle.Solve()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("%v\n", result)
	}
}
