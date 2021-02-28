package day02

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	var tests = []struct {
		passwordLine string
		want         dbEntry
	}{
		{"1-3 a: abcde", dbEntry{1, 3, "a", "abcde"}},
		{"1-3 b: cdefg", dbEntry{1, 3, "b", "cdefg"}},
		{"2-9 c: ccccccccc", dbEntry{2, 9, "c", "ccccccccc"}},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.passwordLine)
		t.Run(testname, func(t *testing.T) {
			ans, _ := parseFile(test.passwordLine)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
