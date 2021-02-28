package day06

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func countUniqueAnswers(answers []string) int {
	count := 0
	uniques := make(map[int32]int)

	for i := 0; i < len(answers); i++ {
		for _, char := range answers[i] {
			if i > 0 {
				if _, ok := uniques[char]; ok {
					uniques[char]++
				}
			} else {
				uniques[char]++
			}
		}
	}

	for _, uniqueAnswer := range uniques {
		if uniqueAnswer == len(answers) {
			count++
		}
	}

	return count
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

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

			answersCount += countUniqueAnswers(answers)

			answers = make([]string, 0)
			continue
		}
		answers = append(answers, line)
	}

	fmt.Printf("Day06 Part2 result: %d\n", answersCount)
}
