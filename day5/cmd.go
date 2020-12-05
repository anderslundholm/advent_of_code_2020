package day5

import "github.com/spf13/cobra"

var (
	part1Cmd = &cobra.Command{
		Use:   "1",
		Short: "run day5-part1",
		Long:  "Run the Day 5, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
	part2Cmd = &cobra.Command{
		Use:   "2",
		Short: "run day5-part2",
		Long:  "Run the Day 5, Part 2 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part2()
		},
	}
)

// AppendCommand adds command for the challange day and part numbers.
func AppendCommand(rootCmd *cobra.Command) {
	dayCmd := &cobra.Command{
		Use:   "day5",
		Short: "Problems for Day 5",
	}
	dayCmd.AddCommand(part1Cmd)
	dayCmd.AddCommand(part2Cmd)

	rootCmd.AddCommand(dayCmd)
}
