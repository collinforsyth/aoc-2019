package cmd

import (
	"fmt"
	"github.com/collinforsyth/aoc-2019/day01"
	"github.com/collinforsyth/aoc-2019/day02"
	"github.com/collinforsyth/aoc-2019/day03"
	"github.com/collinforsyth/aoc-2019/day04"
	"github.com/collinforsyth/aoc-2019/day05"
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
	day02Cmd = &cobra.Command{
		Use: "day02",
		Run: func(cmd *cobra.Command, args []string) {
			day02.Driver()
		},
	}
	day03Cmd = &cobra.Command{
		Use: "day03",
		Run: func(cmd *cobra.Command, args []string) {
			day03.Driver()
		},
	}
	day04Cmd = &cobra.Command{
		Use: "day04",
		Run: func(cmd *cobra.Command, args []string) {
			day04.Driver()
		},
	}
	day05Cmd = &cobra.Command{
		Use: "day05",
		Run: func(cmd *cobra.Command, args []string) {
			day05.Driver()
		},
	}
)

func init() {
	rootCmd.AddCommand(day01Cmd)
	rootCmd.AddCommand(day02Cmd)
	rootCmd.AddCommand(day03Cmd)
	rootCmd.AddCommand(day04Cmd)
	rootCmd.AddCommand(day05Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
