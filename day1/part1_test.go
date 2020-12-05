package day1

import (
	"fmt"
	"testing"
)

func TestMultiplyTwo(t *testing.T) {
	var tests = []struct {
		data []int
		want int
	}{
		{[]int{2019, 1}, 2019},
		{[]int{1000, 10, 10, 1020}, 1020000},
		{[]int{-20, 2040, 15}, -40800},
		{[]int{2000, 2, 3, 20, 5, 5}, 40000},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.data)
		t.Run(testname, func(t *testing.T) {
			ans, _ := multiplyTwo(test.data)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}

// Part1 ...
// func Part1() {
// 	defer timer.ExecutionTimer("Part1")()

// 	ints, err := reader.ReadInts("day1/input.txt")
// 	if err != nil {
// 		log.Fatalf("Could not read ints: %v\n", err)
// 	}

// 	product, err := multiplyTwo(ints)
// 	if err != nil {
// 		log.Fatalf("Could not find value: %v\n", err)
// 	}

// 	fmt.Printf("Day1 Part1 result: %d\n", product)
// }
