package day3

import (
	"fmt"
	"log"
	"os"

	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
	"github.com/spf13/cobra"
)

func part2Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "run day3-part2",
		Long:  "Run the Day 3, Part 2 code.",
		Run: func(cmd *cobra.Command, args []string) {
			Part2()
		},
	}
}

type traversePath struct {
	right int
	down  int
}

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	defer f.Close()

	lines, err := readLines(f)
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	paths := []traversePath{
		traversePath{
			right: 1,
			down:  1,
		},
		traversePath{
			right: 3,
			down:  1,
		},
		traversePath{
			right: 5,
			down:  1,
		},
		traversePath{
			right: 7,
			down:  1,
		},
		traversePath{
			right: 1,
			down:  2,
		},
	}
	result := 1
	for _, path := range paths {
		location := path.right
		treeCount := 0

		for i := path.down; i < len(lines); i += path.down {
			if traverseMap(lines[i], location) {
				treeCount++
			}
			location += path.right
		}
		result = result * treeCount
	}

	fmt.Printf("Day3 Part2 result: %d\n", result)
}