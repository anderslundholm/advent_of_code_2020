package day01

import (
	"fmt"
	"testing"
)

func TestMultiplyThree(t *testing.T) {
	var tests = []struct {
		data []int
		want int
	}{
		{[]int{2018, 1, 1}, 2018},
		{[]int{1000, 10, 10, 40, 980, 1020}, 39200000},
		{[]int{-20, 2040, -30, 10, 15, 1}, -612000},
		{[]int{2000, 2, 3, 20, 5, 15}, 150000},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d", test.data)
		t.Run(testname, func(t *testing.T) {
			ans, _ := multiplyThree(test.data)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
