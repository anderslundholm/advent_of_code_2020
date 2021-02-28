package day05

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day05/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var seatIDs []int
	for _, line := range lines {
		seatID := parseBoardingpass(line)
		seatIDs = append(seatIDs, seatID)
	}

	maxID := 0
	for _, id := range seatIDs {
		if id > maxID {
			maxID = id
		}
	}

	fmt.Printf("Day05 Part1 result: %d\n", maxID)
}
