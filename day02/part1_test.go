package day02

import (
	"fmt"
	"testing"
)

func TestCheckPolicy(t *testing.T) {
	var tests = []struct {
		entry dbEntry
		want  bool
	}{
		{dbEntry{1, 3, "a", "abcde"}, true},
		{dbEntry{1, 3, "b", "cdefg"}, false},
		{dbEntry{2, 9, "c", "ccccccccc"}, true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.entry)
		t.Run(testname, func(t *testing.T) {
			ans := checkPolicy(test.entry)
			if ans != test.want {
				t.Errorf("got %t, want %t", ans, test.want)
			}
		})
	}
}
