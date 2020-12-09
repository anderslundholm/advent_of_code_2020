package day9

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadInts("day9/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	invalidNumber := findInvalidNumber(lines)

	fmt.Printf("Day9 Part1 result: %v\n", invalidNumber)
}
