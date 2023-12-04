package results

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	aoc2021 "github.com/vitorf7/advent-of-code/2021/exercises"
)

var year int

// NewCommand creates a new cobra.Command for the results subcommand
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "results",
		Aliases: []string{"results", "r"},
		Short:   "Command to run and print the Results for the given Advent of Code Year",
		Long:    "Command to run and print the Results for the given Advent of Code Year",
	}

	now := time.Now()
	cmd.Flags().IntVarP(&year, "year", "y", now.Year(), "The year of the Advent of Code to setup a template for")

	cmd.Run = run
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	switch year {
	case 2021:
		aoc2021.PrintResults()
		return
	default:
		fmt.Println(fmt.Sprintf("There are no results for the year %d", year))
		return
	}
}
