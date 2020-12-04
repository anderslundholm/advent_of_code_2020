package day3

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day3/input.txt")
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
