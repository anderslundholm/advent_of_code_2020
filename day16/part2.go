package day16

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input, err := reader.ReadLines("day16/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result, _ := getErrorsV2(input)

	fmt.Printf("Day16 Part2 result: %v\n", result)
}
