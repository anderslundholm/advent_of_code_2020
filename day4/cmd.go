package day4

import "github.com/spf13/cobra"

var (
	part1Cmd = &cobra.Command{
		Use:   "1",
		Short: "run day4-part1",
		Long:  "Run the Day 4, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
	part2Cmd = &cobra.Command{
		Use:   "2",
		Short: "run day4-part2",
		Long:  "Run the Day 4, Part 2 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part2()
		},
	}
)

// AppendCommand adds command for the challange day and part numbers.
func AppendCommand(rootCmd *cobra.Command) {
	dayCmd := &cobra.Command{
		Use:   "day4",
		Short: "Problems for Day 4",
	}
	dayCmd.AddCommand(part1Cmd)
	dayCmd.AddCommand(part2Cmd)

	rootCmd.AddCommand(dayCmd)
}
