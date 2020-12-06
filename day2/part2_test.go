package day2

import (
	"fmt"
	"testing"
)

func TestCheckNewPolicy(t *testing.T) {
	var tests = []struct {
		entry dbEntry
		want  bool
	}{
		{dbEntry{1, 3, "a", "abcde"}, true},
		{dbEntry{1, 3, "b", "cdefg"}, false},
		{dbEntry{2, 9, "c", "ccccccccc"}, false},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.entry)
		t.Run(testname, func(t *testing.T) {
			ans := checkNewPolicy(test.entry)
			if ans != test.want {
				t.Errorf("got %t, want %t", ans, test.want)
			}
		})
	}
}
