package puzzle_helpers

import (
	"strconv"
	"strings"
)

// ConvertMultiLineStringToIntSlice converts a multi-line string to a slice of ints
func ConvertMultiLineStringToIntSlice(input string) ([]int, error) {
	stringSlice := SplitStringToSlice(input, "\n")

	return ConvertStringSliceToInts(stringSlice)
}

// SplitStringToSlice splits a string into a slice of strings
func SplitStringToSlice(inputString string, separator string) []string {
	return strings.Split(inputString, separator)
}

// ConvertStringSliceToInts Helper function to convert a string slice to an integer slice
func ConvertStringSliceToInts(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i, stringValue := range stringSlice {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return nil, err
		}
		intSlice[i] = intValue
	}

	return intSlice, nil
}
