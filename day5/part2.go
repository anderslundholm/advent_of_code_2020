package day5

import (
	"fmt"
	"log"
	"sort"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day5/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var seatIDs []int
	for _, line := range lines {
		seatID := parseBoardingpass(line)
		seatIDs = append(seatIDs, seatID)
	}

	var mySeat int
	sort.Ints(seatIDs)
	for i := 0; i < len(seatIDs)-1; i++ {
		// The seats with IDs +1 and -1 from yours will be in your list.
		// Check if the next expected seatID is equal to the next seatID in the list.
		if seatIDs[i]+1 != seatIDs[i+1] {
			mySeat = seatIDs[i] + 1
		}
	}

	fmt.Printf("Day5 Part2 result: %d\n", mySeat)
}
