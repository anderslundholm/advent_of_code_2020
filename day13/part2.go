package day13

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"

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

type bus struct {
	id     int
	offset int
}

func earliestBusOffset(input []string) int {
	var bussIDs []int
	var bussIDIndexes = make(map[int]int)
	var busses []bus
	departTimes := strings.Split(input[1], ",")
	for i, t := range departTimes {
		if t != "x" {
			bussIDs = append(bussIDs, util.MustAtoi(t))
			bussIDIndexes[util.MustAtoi(t)] = i

			busses = append(busses, bus{id: util.MustAtoi(t), offset: i})
			// fmt.Println(i, util.MustAtoi(t))
		}
	}
	// increment := 1
	// time := 1
	// for bussTime, offset := range bussIDIndexes {

	// 	for {
	// 		nextBussTime := time + offset
	// 		if nextBussTime%bussTime == 0 {
	// 			break
	// 		}
	// 		time += increment

	// 	}
	// 	increment *= offset
	// 	fmt.Println(time, increment, offset)
	// }

	// b := 77
	// offset := 1
	// inc := 1
	for i := 0; i < len(busses)-1; i++ {
		nextBuss := busses[i+1].id
		for j := busses[i].id; ; j = j + busses[i].id {

			if nextBuss <= j {
				nextBuss += busses[i+1].id
			}
			// for k := busses[i+1].id; ; k = k + busses[i+1].id {

			// }
			if j == nextBuss-1 {
				// fmt.Println(j, nextBuss)
				break
			}

			fmt.Println(j, nextBuss)
			time.Sleep(1 * time.Second)
		}
		// a := util.LCM(bussIDs[i], bussIDs[i+1])
		// a + b
		// fmt.Println(busses[i])

	}
	result := 0
	return result
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input, err := reader.ReadLines("day13/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := earliestBusOffset(input)

	fmt.Printf("Day13 Part2 result: %v\n", result)
}
