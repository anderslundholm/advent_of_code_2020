package day13

import (
	"fmt"
	"log"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

func earliestBusToAirport(input []string) int {
	var earliestBus int
	var earliestBusTime int
	earliestTimestamp := util.MustConvertAtoi(input[0])

	departTimes := strings.Split(input[1], ",")
	for _, t := range departTimes {
		if t != "x" {
			t := util.MustConvertAtoi(t)
			mod := earliestTimestamp % t
			if earliestBusTime == 0 || earliestBusTime > t-mod {
				earliestBus = t
				earliestBusTime = t - mod
			}
		}
	}
	return earliestBus * earliestBusTime
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	input, err := reader.ReadLines("day13/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := earliestBusToAirport(input)

	fmt.Printf("Day13 Part1 result: %v\n", result)
}
