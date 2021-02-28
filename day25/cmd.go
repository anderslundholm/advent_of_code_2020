package day25

import "github.com/spf13/cobra"

type partFunc func()

func getPartCmd(part string, day string, f partFunc) *cobra.Command {
	return &cobra.Command{
		Use:   part,
		Short: "run day" + day + "-part" + part,
		Long:  "Run the Day " + day + ", Part " + part + " code.",
		Run: func(cmd *cobra.Command, args []string) {
			f()
		},
	}
}

// AppendCommand adds command for the challange day and part numbers.
func AppendCommand(rootCmd *cobra.Command) {
	day := "25"
	dayCmd := &cobra.Command{
		Use:   "day" + day,
		Short: "Problems for Day " + day,
	}
	part1Cmd := getPartCmd("1", day, Part1)
	dayCmd.AddCommand(part1Cmd)

	rootCmd.AddCommand(dayCmd)
}
