package day13

import (
	"fmt"
	"log"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Testing implementing CRT
func chineseRemainder() {
	// The Chinese remainder theorem asserts that if the ni are pairwise coprime, and if a1, ..., ak
	// are integers such that 0 ≤ ai < ni for every i, then there is one and only one integer x, such
	// that 0 ≤ x < N and the remainder of the Euclidean division of x by ni is ai for every i.
}

type bus struct {
	id     int
	offset int
}

// Testing implementing with LCM (broken or just very slow).
func leastCommonMultiple(input []string) int {
	var bussIDs []int
	var bussIDIndexes = make(map[int]int)
	var busses []bus
	departTimes := strings.Split(input[1], ",")
	for i, t := range departTimes {
		if t != "x" {
			bussIDs = append(bussIDs, util.MustAtoi(t))
			bussIDIndexes[util.MustAtoi(t)] = i
			busses = append(busses, bus{id: util.MustAtoi(t), offset: i})
		}
	}

	var lastBus int
	for i := 0; i < len(busses)-1; i++ {
		curBus := busses[i].id
		nextBus := busses[i+1].id
		increment := curBus

		if i == 1 {
			increment = util.LCM(busses[0].id, busses[1].id)
		} else if i > 1 {
			increment = util.LCM(busses[0].id, busses[1].id, bussIDs[2:i+1]...)
		} else {
			lastBus = curBus
		}
		fmt.Println(lastBus, curBus, nextBus, increment, busses[i+1].offset, bussIDs[i:i+1])
		for j := lastBus; ; j = j + increment {
			for {
				if nextBus <= j {
					nextBus += busses[i+1].id
				} else {
					break
				}
			}
			if j == nextBus-busses[i+1].offset {
				lastBus = j
				break
			}
		}
	}
	return lastBus
}

func earliestBusAlignment(input []string) int {
	var bussIDIndexes = make(map[int]int)
	departTime := 1
	increment := 1
	departTimes := strings.Split(input[1], ",")
	for i, t := range departTimes {
		if t != "x" {
			bussIDIndexes[util.MustAtoi(t)] = i
		}
	}

	for busID, offset := range bussIDIndexes {
		for {
			if (departTime+offset)%busID == 0 {
				break
			}
			departTime += increment
		}
		increment *= busID
		fmt.Println(busID, offset, departTime, increment)
	}
	return departTime
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	input, err := reader.ReadLines("day13/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := earliestBusAlignment(input)

	fmt.Printf("Day13 Part2 result: %v\n", result)
}
