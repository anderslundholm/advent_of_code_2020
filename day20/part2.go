package day20

import (
	"fmt"
	"io/ioutil"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	// input, err := reader.ReadLines("day20/input.txt")
	// if err != nil {
	// 	log.Fatalf("Could not read lines: %v\n", err)
	// }

	input, _ := ioutil.ReadFile("day20/input.txt")

	result := part2(string(input))

	fmt.Printf("Day20 Part2 result: %v\n", result)
}
