package day6

import (
	"fmt"
	"testing"
)

func TestCountAnswers(t *testing.T) {
	var tests = []struct {
		answers []string
		want    int
	}{
		{[]string{"a", "b"}, 2},
		{[]string{"a", "abc", "xyb", "asdfb", "a"}, 8},
		{[]string{"a"}, 1},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.answers)
		t.Run(testname, func(t *testing.T) {
			ans := countAnswers(test.answers)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
