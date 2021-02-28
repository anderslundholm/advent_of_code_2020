package day02

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func checkPolicy(entry dbEntry) bool {
	charInstance := 0
	for _, char := range entry.password {
		if string(char) == entry.character {
			charInstance++
		}
	}

	return charInstance >= entry.min && charInstance <= entry.max
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

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
		if checkPolicy(entry) {
			count++
		}
	}

	fmt.Printf("Day02 Part1 result: %d\n", count)
}
