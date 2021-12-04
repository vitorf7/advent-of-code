package setuptemplate

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	downloadinput "github.com/vitorf7/advent-of-code/cmd/download_input"
)

var (
	//go:embed puzzle.go.tmpl
	puzzleTemplate string

	//go:embed puzzle_test.go.tmpl
	puzzleTestTemplate string

	year          int
	day           int
	downloadInput bool
)

type setup struct {
	Year int
	Day  int
}

// NewCommand creates a new cobra.Command for the setup-template subcommand
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setup-template",
		Aliases: []string{"setup", "st"},
		Short:   "Command to setup a base template for the given Advent of Code Day",
		Long:    "Command to setup a base template for the given Advent of Code Day",
	}

	now := time.Now()
	cmd.Flags().IntVarP(&day, "day", "d", now.Day(), "The day of the Advent of Code to setup a template for")
	cmd.Flags().IntVarP(&year, "year", "y", now.Year(), "The year of the Advent of Code to setup a template for")
	cmd.Flags().BoolVarP(&downloadInput, "download-input", "i", false, "Download the input file for the Advent of Code Day")

	cmd.Run = run
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	setup := setup{Year: year, Day: day}

	newDirPath := fmt.Sprintf("%d/day%d", setup.Year, setup.Day)
	err := os.MkdirAll(newDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(fmt.Errorf("MkdirAll %q: %s ", newDirPath, err).Error())
	}

	// Create input.txt
	inputFilePath := fmt.Sprintf("%s/input.txt", newDirPath)
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		panic(err)
	}

	// Create puzzle.go
	err = createFile(newDirPath+"/puzzle.go", puzzleTemplate, setup)
	if err != nil {
		panic(err)
	}

	// Create puzzle_test.go
	err = createFile(newDirPath+"/puzzle_test.go", puzzleTestTemplate, setup)
	if err != nil {
		panic(err)
	}

	if downloadInput {
		if err = downloadinput.DownloadInput(setup.Year, setup.Day, inputFile); err != nil {
			panic(err)
		}
	}
}

func createFile(path string, templateFile string, setup setup) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}

	tmpl, err := template.New("setup").Parse(templateFile)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, setup)
	if err != nil {
		return err
	}

	return nil
}
