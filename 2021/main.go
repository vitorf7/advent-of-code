package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	downloadinput "github.com/vitorf7/advent-of-code/2021/cmd/download_input"
	"github.com/vitorf7/advent-of-code/2021/cmd/results"
	setuptemplate "github.com/vitorf7/advent-of-code/2021/cmd/setup_template"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	root := &cobra.Command{
		Use:   "advent-of-code",
		Short: "Advent of Code",
		Long:  "A repo to solve Advent of Code problems, download inputs and create base working files for the solutions",
	}

	root.AddCommand(setuptemplate.NewCommand())
	root.AddCommand(downloadinput.NewCommand())
	root.AddCommand(results.NewCommand())

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
