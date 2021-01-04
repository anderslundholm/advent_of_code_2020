package day21

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	input, err := reader.ReadLines("day21/test_input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	result := getAlergenFreeIngredients(input)

	fmt.Printf("Day21 Part1 result: %v\n", result)
}
