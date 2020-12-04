package day1

import (
	"errors"
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part2Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "run day1-part2",
		Long:  "Run the Day 1, Part 2 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part2()
		},
	}
}

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

	ints, err := reader.ReadInts("day1/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}
	product, err := multiplyThree(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day1 Part1 result: %d\n", product)
}
