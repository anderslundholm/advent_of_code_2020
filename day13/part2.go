package day13

import (
	"fmt"
	"log"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// The Chinese remainder theorem asserts that if the ni are pairwise coprime, and if a1, ..., ak
// are integers such that 0 ≤ ai < ni for every i, then there is one and only one integer x, such
// that 0 ≤ x < N and the remainder of the Euclidean division of x by ni is ai for every i.
// func chineseRemainder() {

// }

// func earliestBusOffset(input []string) int {
// 	var earliestBus int
// 	var earliestBusTime int

// 	departTimes := strings.Split(input[1], ",")
// 	fmt.Println(len(departTimes))
// 	for _, t := range departTimes {
// 		if t == "x" {

// 		}
// 	}
// 	return earliestBus * earliestBusTime
// }

func test(input []string) int {
	var result int
	// var bussIDs []int
	departTimes := strings.Split(input[1], ",")
	for _, t := range departTimes {
		if t != "x" {
			// bussIDs = t
		}
	}

	return result
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input, err := reader.ReadLines("day13/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := test(input)

	fmt.Printf("Day13 Part2 result: %v\n", result)
}
