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
	last  int
	rule  int
	match bool
}

func check(ruleResult memoizeRule, rules map[int]rule, char string) memoizeRule {
	// fmt.Println(string(char))
	ruleResult.match = false
	if rules[ruleResult.rule].char != "" {
		fmt.Println("hell?", ruleResult.rule, char, ruleResult)
		// fmt.Println("asdf", ruleNum, rules[ruleNum].char, char)
		if char == rules[ruleResult.rule].char {
			// fmt.Println(ruleNum, char)
			ruleResult.match = true
			ruleResult.last++
			return ruleResult
		}
		ruleResult.match = false
		return ruleResult
	} else {
		curRuleNum := rules[ruleResult.rule].rules0
		for i := ruleResult.last; i < len(curRuleNum); i++ {
			ruleResult.last = i
			ruleResult.rule = curRuleNum[i]
			ruleResult = check(ruleResult, rules, char)
			fmt.Println(ruleResult.rule, curRuleNum[i], char, ruleResult)
			if !ruleResult.match {
				break
			} else {
				return ruleResult
			}
		}
		if ruleResult.match {
			return ruleResult
		}
		curRuleNum = rules[ruleResult.rule].rules0
		for i := ruleResult.last; i < len(curRuleNum); i++ {
			ruleResult.last = i
			ruleResult.rule = curRuleNum[i]
			ruleResult = check(ruleResult, rules, char)
			fmt.Println(ruleResult.rule, curRuleNum[i], char, ruleResult)
			if !ruleResult.match {
				break
			} else {
				return ruleResult
			}
		}
	}
	return ruleResult
}

func match(message string, start, end int, rules map[int]rule, r int) bool {
	// memoize := memoizeRule{start: start, end: end, rule: r}
	// if ok := discoveredMessages[memoize]; ok {
	// 	return true
	// }
	// fmt.Println(message)
	var ruleResult = memoizeRule{rule: 0, last: 0, match: false}
	for _, char := range message {
		fmt.Println("Checking: ", message, ruleResult)
		ruleResult = check(ruleResult, rules, string(char))
		if !ruleResult.match {
			fmt.Println("!!!!", message, ruleResult)
			return false
		}

	}
	fmt.Println("returned true", message, ruleResult)

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
			fmt.Println("result", result)
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
