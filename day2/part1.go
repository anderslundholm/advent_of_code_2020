package day2

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part1Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "run day2-part1",
		Long:  "Run the Day 2, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
}

func checkPolicy(entry dbEntry) bool {
	charInstance := 0
	for _, char := range entry.password {
		if string(char) == entry.character {
			charInstance++
		}
	}

	return charInstance >= entry.min && charInstance <= entry.max
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day2/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}
	count := 0

	for _, line := range lines {
		entry, err := parseFile(line)
		if err != nil {
			log.Printf("Could not parse entry: %v", err)
		}
		if checkPolicy(entry) {
			count++
		}
	}

	fmt.Printf("Day2 Part1 result: %d\n", count)
}
