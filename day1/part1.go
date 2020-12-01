package day1

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part1Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "run day1-part1",
		Long:  "Run the Day 1, Part 1 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part1()
		},
	}
}

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

	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	defer f.Close()

	ints, err := readInts(f)

	product, err := multiplyTwo(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day1 Part1 result: %d\n", product)
}
