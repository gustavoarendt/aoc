package cmd

import (
	"fmt"
	"os"

	"github.com/gustavoarendt/aoc/cmd/2023/day01"
	"github.com/gustavoarendt/aoc/cmd/2023/day02"
	"github.com/gustavoarendt/aoc/cmd/2023/day03"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "Advent of Code",
	Long:  "Advent of Code using Go Language",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(day01.Cmd)
	Cmd.AddCommand(day02.Cmd)
	Cmd.AddCommand(day03.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Help() {
	print("Uses go run main.go 'year' 'day' to run the program\n")
}
