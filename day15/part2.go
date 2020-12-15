package day15

import (
	"fmt"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input := []int{6, 13, 1, 15, 2, 0}

	result := getLastNumberV2(input, 30000000)

	fmt.Printf("Day15 Part2 result: %v\n", result)
}
