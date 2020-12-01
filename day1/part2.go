package day1

import (
	"errors"
	"fmt"
	"log"
	"os"

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

	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	defer f.Close()

	ints, err := readInts(f)

	product, err := multiplyThree(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day1 Part1 result: %d\n", product)
}
