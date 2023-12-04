package downloadinput

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	year int
	day  int
)

// NewCommand creates a new cobra.Command for the download-input subcommand
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download-input",
		Aliases: []string{"download", "di"},
		Short:   "Command to download input for the given Advent of Code Day",
		Long:    "Command to download input for the given Advent of Code Day",
	}

	now := time.Now()
	cmd.Flags().IntVarP(&day, "day", "d", now.Day(), "The day of the Advent of Code to setup a template for")
	cmd.Flags().IntVarP(&year, "year", "y", now.Year(), "The year of the Advent of Code to setup a template for")

	cmd.Run = run
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	inputFile, err := os.OpenFile(fmt.Sprintf("%d/day%d/input.txt", year, day), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}

	err = DownloadInput(year, day, inputFile)
}

func DownloadInput(year int, day int, targetFile *os.File) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	client := http.Client{
		Jar: jar,
	}

	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    os.Getenv("ADVENT_OF_CODE_COOKIE"),
		Path:     "/",
		Domain:   ".adventofcode.com",
		Secure:   true,
		HttpOnly: true,
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.AddCookie(sessionCookie)

	// Get the data
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(targetFile, resp.Body)
	return err
}
