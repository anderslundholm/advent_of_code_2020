package day19

import (
	"fmt"
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

var discoveredMessages = make(map[memoizeRule]bool)

type memoizeRule struct {
	start int
	end   int
	rule  int
}

func check(ruleNum int, rules map[int]rule, char string, i int) bool {
	// fmt.Println(string(char))
	if rules[ruleNum].char != "" {
		fmt.Println(ruleNum, char)
		if char == rules[ruleNum].char {
			fmt.Println(ruleNum, char)
			return true
		}
		return false
	} else {
		result := false
		for _, num := range rules[ruleNum].rules0 {
			fmt.Println(ruleNum, num)
			result = check(num, rules, char)
			if !result {
				break
			}
		}
		if result {
			return true
		}
		for _, num := range rules[ruleNum].rules1 {
			result = check(num, rules, char)
			if !result {
				break
			}
		}
		if result {
			return true
		}
	}
	return false
}

func match(message string, start, end int, rules map[int]rule, r int) bool {
	memoize := memoizeRule{start: start, end: end, rule: r}
	if ok := discoveredMessages[memoize]; ok {
		return true
	}
	fmt.Println(message)
	for i, char := range message {

		if !check(0, rules, string(char), i) {
			fmt.Println("!!!!", message)
			return false
		}
	}
	// fmt.Println(message)
	return true
}

func matchingRules(input []string) int {
	result := 0
	rules, messages := parseInput(input)

	fmt.Println(rules[0], rules)
	// recurseRules(rules, 0)
	// for _, rule := range rules {
	// 	recurseRules(rule)
	// }
	for message := range messages {
		// fmt.Println(message)
		if match(message, 0, len(message), rules, 0) {

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
