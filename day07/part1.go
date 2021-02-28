package day07

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day07/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var allLines []string
	for _, line := range lines {
		allLines = append(allLines, line)

	}

	rules := parseBags(allLines)
	result := 0
	for rule := range rules {
		if rules.checkIfBagContainsBag(rule) {
			result++
		}
	}

	fmt.Printf("Day07 Part1 result: %v\n", result)
}
