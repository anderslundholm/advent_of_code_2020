package day4

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

type traversePath struct {
	right int
	down  int
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day4/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	passport := make([]string, 0)
	var validCount int
	for i, line := range lines {
		if line == "" {
			p := parsePassport(passport)
			if p.isValidFields() {
				validCount++
			}
			passport = make([]string, 0)
			continue
		} else if i == len(lines)-1 {
			passport = append(passport, line)
			p := parsePassport(passport)
			if p.isValidFields() {
				validCount++
			}

		}
		passport = append(passport, line)
	}

	fmt.Printf("Day4 Part2 result: %d\n", validCount)
}
