package day01

import (
	"errors"
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func multiplyThree(data []int) (int, error) {
	var product int
	err := errors.New("no value")
	for _, x := range data {
		for _, y := range data {
			for _, z := range data {
				if sum := x + y + z; sum == 2020 {
					product = x * y * z
					return product, nil
				}
			}
		}
	}
	return product, err
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	ints, err := reader.ReadInts("day01/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}
	product, err := multiplyThree(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day01 Part1 result: %d\n", product)
}
