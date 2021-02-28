package day01

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
