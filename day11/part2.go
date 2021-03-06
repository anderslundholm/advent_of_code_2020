package day11

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func modelNewProcess(input []int) {

}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input, err := reader.ReadMultiLineRunes("day11/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}
	// result := input
	result := newModelProcess(&input)

	fmt.Printf("Day11 Part2 result: %v\n", result)
}
