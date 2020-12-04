package day4

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
		Short: "run day4-part1",
		Long:  "Run the Day 2, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day4/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	passport := make([]string, 0)
	var validCount int
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			if i == len(lines)-1 {
				passport = append(passport, line)
			}
			p := parsePassport(passport)
			if p.isValid() {
				validCount++
			}
			passport = make([]string, 0)
			continue
		}
		passport = append(passport, line)
	}
	fmt.Printf("Day4 Part1 result: %v\n", validCount)
}
