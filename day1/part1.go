package day1

import (
	"errors"
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func multiplyTwo(data []int) (int, error) {
	var product int
	err := errors.New("no value")
	for _, x := range data {
		for _, y := range data {
			if sum := x + y; sum == 2020 {
				product = x * y
				return product, nil
			}
		}
	}
	return product, err
}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	ints, err := reader.ReadInts("day1/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	product, err := multiplyTwo(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day1 Part1 result: %d\n", product)
}
