package day3

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

	lines, err := reader.ReadLines("day3/input.txt")
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
