package day10

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadInts("day10/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := countChargerArrangements(lines)

	fmt.Printf("Day10 Part2 result: %v\n", result)
}
