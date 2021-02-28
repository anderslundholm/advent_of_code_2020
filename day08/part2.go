package day08

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day08/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var allLines []string
	for _, line := range lines {
		allLines = append(allLines, line)

	}
	// visitedLines := make(map[int]bool)
	// changedInstructions := make(map[int]bool)
	// result, _ := fixLoop(allLines, 0, 0, false, visitedLines, changedInstructions)

	result := runProgram2(lines)

	fmt.Printf("Day08 Part2 result: %v\n", result)
}
