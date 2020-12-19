package day19

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	input, err := reader.ReadLines("day19/test_input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := matchingRules(input)

	fmt.Printf("Day19 Part1 result: %v\n", result)
}
