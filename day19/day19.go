package day19

import (
	"sort"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

type rule struct {
	index  int
	rules0 []int
	rules1 []int
	char   string
}

func (r *rule) addRule(i int, n int) *rule {
	if n == 0 {
		r.rules0 = append(r.rules0, i)
	} else if n == 1 {
		r.rules1 = append(r.rules1, i)
	}
	return r
}

type recurseResult struct {
	ruleIndex int
	index     int
	next      []int
}

func matchingRules(input []string, part2 bool) int {
	var result int
	rules, messages := parseInput(input)
	rr := recurseResult{ruleIndex: 0, index: 0, next: nil}

	if part2 {
		rules[8] = rule{index: 8, rules0: []int{42}, rules1: []int{42, 8}}
		rules[11] = rule{index: 11, rules0: []int{42, 31}, rules1: []int{42, 11, 31}}
	}

	for message := range messages {
		if check(message, rules, rr) {
			result++
		}
	}
	return result
}

func parseInput(input []string) (rules map[int]rule, messages map[string]bool) {
	sort.Strings(input)
	rules = make(map[int]rule)
	messages = make(map[string]bool)
	for _, line := range input {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, ":")
		switch len(splitLine) {
		case 1:
			messages[line] = true
		case 2:
			r := rule{}
			index := util.MustAtoi(splitLine[0])
			if strings.Contains(splitLine[1], "\"") {
				rules[index] = rule{index: index, char: strings.Trim(splitLine[1], ` "`)}
			} else {
				multiSplitLine := strings.Split(splitLine[1], " | ")
				for i := 0; i < len(multiSplitLine); i++ {
					for _, stringField := range strings.Fields(multiSplitLine[i]) {
						r.addRule(util.MustAtoi(stringField), i)
						rules[index] = r
					}
				}
			}
		}
	}
	return rules, messages
}

func check(message string, rules map[int]rule, rr recurseResult) bool {
	rule := rules[rr.ruleIndex]
	if rule.char != "" {
		if string(message[rr.index]) != rule.char {
			return false
		}
		if len(rr.next) == 0 {
			return rr.index == len(message)-1
		} else if rr.index+1 >= len(message) {
			return false
		} else {
			rr2 := recurseResult{ruleIndex: rr.next[0], index: rr.index + 1, next: rr.next[1:]}
			return check(message, rules, rr2)
		}
	}

	nextRule := rule.rules0[1:]
	next := make([]int, len(rr.next)+len(nextRule))
	copy(next, nextRule)
	copy(next[len(nextRule):], rr.next)
	rr2 := recurseResult{ruleIndex: rule.rules0[0], index: rr.index, next: next}
	if check(message, rules, rr2) {
		return true
	}

	if len(rule.rules1) != 0 {
		nextRule := rule.rules1[1:]
		next := make([]int, len(rr.next)+len(nextRule))
		copy(next, nextRule)
		copy(next[len(nextRule):], rr.next)
		rr2 := recurseResult{ruleIndex: rule.rules1[0], index: rr.index, next: next}
		if check(message, rules, rr2) {
			return true
		}
	}
	return false
}
