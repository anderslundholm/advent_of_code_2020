package day07

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseBags(t *testing.T) {
	var tests = []struct {
		lines []string
		want  bagRules
	}{
		{[]string{
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"faded blue bags contain no other bags."},

			map[bagType]map[bagType]int{
				{description: "dark", color: "olive"}: {
					{description: "faded", color: "blue"}:   3,
					{description: "dotted", color: "black"}: 4},
				{description: "faded", color: "blue"}: {}},
		},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.lines)
		t.Run(testname, func(t *testing.T) {
			ans := parseBags(test.lines)
			if reflect.DeepEqual(ans, test.want) != true {
				t.Errorf("got %v want %v", ans, test.want)
			}
		})
	}
}
