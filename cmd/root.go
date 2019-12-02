package cmd

import (
	"fmt"
	"github.com/collinforsyth/aoc-2019/day01"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 Driver Program",
		Args:  cobra.MinimumNArgs(1),
	}
	day01Cmd = &cobra.Command{
		Use: "day01",
		Run: func(cmd *cobra.Command, args []string) {
			day01.Driver()
		},
	}
)

func init() {
	rootCmd.AddCommand(day01Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}