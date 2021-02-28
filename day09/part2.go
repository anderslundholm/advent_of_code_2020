package day09

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadInts("day09/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	invalidNumber := findInvalidNumber(lines)
	encryptionWeakness := findEncryptionWeakness(lines, invalidNumber)

	fmt.Printf("Day09 Part2 result: %v\n", encryptionWeakness)
}
