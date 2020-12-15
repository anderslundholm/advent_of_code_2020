package day15

import (
	"fmt"
	"testing"
)

func TestFindInSlice(t *testing.T) {
	var tests = []struct {
		list []int
		i    int
		want int
	}{
		{[]int{1, 3, 2}, 1, 0},
		{[]int{1, 3, 2, 0, 1}, 4, 4},
		{[]int{1, 2, 3, 3}, 3, 1},
		{[]int{2, 3, 1}, 0, 0},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d %d", test.list, test.i)
		t.Run(testname, func(t *testing.T) {
			ans := findInSlice(test.list, test.i)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}

func TestGetLastNumber(t *testing.T) {
	var tests = []struct {
		list  []int
		turns int
		want  int
	}{
		{[]int{1, 3, 2}, 2020, 1},
		{[]int{2, 1, 3}, 2020, 10},
		{[]int{1, 2, 3}, 2020, 27},
		{[]int{2, 3, 1}, 2020, 78},
		{[]int{3, 2, 1}, 2020, 438},
		{[]int{3, 1, 2}, 2020, 1836},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d %d", test.list, test.turns)
		t.Run(testname, func(t *testing.T) {
			ans := getLastNumber(test.list, test.turns)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}

func TestGetLastNumberV2(t *testing.T) {
	var tests = []struct {
		list  []int
		turns int
		want  int
	}{
		{[]int{1, 3, 2}, 30000000, 2578},
		{[]int{2, 1, 3}, 30000000, 3544142},
		{[]int{1, 2, 3}, 30000000, 261214},
		{[]int{2, 3, 1}, 30000000, 6895259},
		{[]int{3, 2, 1}, 30000000, 18},
		{[]int{3, 1, 2}, 30000000, 362},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d %d", test.list, test.turns)
		t.Run(testname, func(t *testing.T) {
			ans := getLastNumberV2(test.list, test.turns)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
