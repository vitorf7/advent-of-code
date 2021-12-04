package results

import "github.com/spf13/cobra"

// NewCommand creates a new cobra.Command for the results subcommand
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "results",
		Aliases: []string{"results", "r"},
		Short:   "Command to run and print the Results for the given Advent of Code Year",
		Long:    "Command to run and print the Results for the given Advent of Code Year",
	}
	cmd.Run = run
	return cmd
}

func run(cmd *cobra.Command, args []string) {}
