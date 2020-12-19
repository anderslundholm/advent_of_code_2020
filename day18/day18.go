package day18

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

func getSumOfResults(input []string) int {
	result := 0
	re := regexp.MustCompile(`\([^\(\)]+\)`)
	for _, line := range input {
		for re.MatchString(line) {
			line = re.ReplaceAllStringFunc(line, func(line string) string { return fmt.Sprint(evaluateExpression(line)) })
		}
		result += evaluateExpression(line)
	}
	return result
}

func getSumOfResultsV2(input []string) int {
	result := 0
	for _, line := range input {
		result += recurseMatch(line, regexp.MustCompile(`\([^\(\)]+\)`), matchPlusOperations)
	}
	return result
}

func matchPlusOperations(line string) int {
	return recurseMatch(line, regexp.MustCompile(`\d+ \+ \d+`), evaluateExpression)
}

func recurseMatch(line string, re *regexp.Regexp, function func(string) int) int {
	for re.MatchString(line) {
		line = re.ReplaceAllStringFunc(line, func(line string) string {
			return fmt.Sprint(function(line))
		})
	}
	return function(line)
}

func evaluateExpression(line string) int {
	chars := strings.Fields(strings.Trim(line, "()"))
	num := util.MustAtoi(chars[0])

	for i := 1; i < len(chars); i += 2 {
		switch chars[i] {
		case "+":
			num += util.MustAtoi(chars[i+1])
		case "*":
			num *= util.MustAtoi(chars[i+1])
		}
	}
	return num
}
