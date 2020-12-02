package day2

import "github.com/spf13/cobra"

// AppendCommand adds command for the challange day and part numbers.
func AppendCommand(rootCmd *cobra.Command) {
	dayCmd := &cobra.Command{
		Use:   "day2",
		Short: "Problems for Day 2",
	}
	dayCmd.AddCommand(part1Cmd())
	dayCmd.AddCommand(part2Cmd())

	rootCmd.AddCommand(dayCmd)
}
