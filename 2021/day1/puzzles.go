package aoc2021day1

import (
	"bufio"
	"strconv"
	"strings"
)

func Puzzle1(puzzleInput string) (int, error) {
	increases := 0
	previousNumber := 0

	inputScanner := bufio.NewScanner(strings.NewReader(puzzleInput))
	for inputScanner.Scan() {
		currentNumber, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			panic(err)
		}

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

func Puzzle2(puzzleInput string, windowSize int) (int, error) {
	increases := 0
	previousNumber := 0

	inputSlice := strings.Split(puzzleInput, "\n")
	start := 0
	currentSum := 0
	for index, valueString := range inputSlice {
		valueInt, err := strconv.Atoi(valueString)
		if err != nil {
			return 0, err
		}

		currentSum += valueInt

		if previousNumber != 0 && currentSum > previousNumber {
			increases++
		}

		if index-start+1 == windowSize {
			previousNumber = currentSum
			startValue, err := strconv.Atoi(inputSlice[start])
			if err != nil {
				return 0, err
			}
			currentSum -= startValue
			start += 1
		}
	}

	return increases, nil
}
