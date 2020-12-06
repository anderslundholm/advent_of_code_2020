package day3

import (
	"fmt"
	"testing"
)

func TestTraverseMap(t *testing.T) {
	var tests = []struct {
		line     string
		location int
		want     bool
	}{
		{"#...#...#..#...#....#...#..", 3, false},
		{".#....#..#..#....#..#..#....#..#..#....#############.", 6, true},
		{"..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#", 9, false},
		{".#...##..#..#...##..#..#...##..#.##..#.", 12, true},
		{"..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....", 15, true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s %d", test.line, test.location)
		t.Run(testname, func(t *testing.T) {
			ans := traverseMap(test.line, test.location)
			if ans != test.want {
				t.Errorf("got %t, want %t", ans, test.want)
			}
		})
	}
}
