package day3

import (
	"fmt"
	"log"
	"os"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part1Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "run day3-part1",
		Long:  "Run the Day 2, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	defer f.Close()

	lines, err := readLines(f)
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	location := 3
	treeCount := 0

	for _, line := range lines[1:] {
		if traverseMap(line, location) {
			treeCount++
		}
		location += 3
	}

	fmt.Printf("Day3 Part1 result: %d\n", treeCount)
}
