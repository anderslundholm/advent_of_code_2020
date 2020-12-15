package day15

import (
	"fmt"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	input := []int{6, 13, 1, 15, 2, 0}

	result := getLastNumber(input, 2020)

	fmt.Printf("Day15 Part1 result: %v\n", result)
}
