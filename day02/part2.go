package day02

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func checkNewPolicy(entry dbEntry) bool {
	return (string(entry.password[entry.min-1]) == entry.character) != (string(entry.password[entry.max-1]) == entry.character)
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day02/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}
	count := 0

	for _, line := range lines {
		entry, err := parseFile(line)
		if err != nil {
			log.Printf("Could not parse entry: %v", err)
		}
		if checkNewPolicy(entry) {
			count++
		}
	}

	fmt.Printf("Day02 Part2 result: %d\n", count)
}
