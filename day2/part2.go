package day2

import (
	"fmt"
	"log"
	"os"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part2Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "run day2-part2",
		Long:  "Run the Day 2, Part 2 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part2()
		},
	}
}

func checkNewPolicy(entry dbEntry) bool {
	return (string(entry.password[entry.min-1]) == entry.character) != (string(entry.password[entry.max-1]) == entry.character)
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	f, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	defer f.Close()

	lines, err := readLines(f)
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}
	count := 0

	for _, line := range lines {
		entry, err := parseFile(line)
		if err != nil {
			log.Printf("Could not parse entry: %v", err)
		}
		if checkNewPolicy(entry) {
			count++
		}
	}

	fmt.Printf("Day2 Part2 result: %d\n", count)
}
