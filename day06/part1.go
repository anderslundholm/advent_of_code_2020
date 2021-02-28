package day06

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func countAnswers(answers []string) int {
	unique := make(map[int32]bool)

	for _, line := range answers {
		for _, char := range line {
			unique[char] = true
		}
	}

	return len(unique)
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day06/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var answers []string
	var answersCount int
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			if i == len(lines)-1 {
				answers = append(answers, line)
			}

			answersCount += countAnswers(answers)

			answers = make([]string, 0)
			continue
		}
		answers = append(answers, line)
	}

	fmt.Printf("Day06 Part1 result: %d\n", answersCount)
}
