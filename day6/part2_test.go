package day6

import (
	"fmt"
	"testing"
)

func TestCountUniqueAnswers(t *testing.T) {
	var tests = []struct {
		answers []string
		want    int
	}{
		{[]string{"abc", "ab"}, 2},
		{[]string{"abc"}, 3},
		{[]string{"abc", "abc"}, 3},
		{[]string{"a", "b", "c"}, 0},
		{[]string{"a", "abc", "axyb", "asdfb", "a"}, 1},
		{[]string{"abc", "abx", "xba"}, 2},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.answers)
		t.Run(testname, func(t *testing.T) {
			ans := countUniqueAnswers(test.answers)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
