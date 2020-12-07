package day7

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day7/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var allLines []string
	for _, line := range lines {
		allLines = append(allLines, line)

	}

	rules := parseBags(allLines)
	result := rules.countBags(myBag) - 1

	fmt.Printf("Day7 Part2 result: %v\n", result)
}