package day16

import (
	"fmt"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

type rule struct {
	index      int
	name       string
	a, b, c, d int
}

func getErrorRate(input []string) int {
	result, _ := getErrors(input)
	return result
}

func getErrors(input []string) (int, []int) {
	var result int
	var yourTicket []int
	var stopPoint int
	var errors []int
	ticketRules := make(map[int]bool)
	var nearbyTickets []int
	for i, line := range input {
		if input[i] == "" {
			break
		}
		splitLine := strings.Split(line, ": ")
		numRange := strings.Split(splitLine[1], " or ")
		numRange1 := strings.Split(numRange[0], "-")
		numRange2 := strings.Split(numRange[1], "-")
		for i := util.MustAtoi(numRange1[0]); i <= util.MustAtoi(numRange1[1]); i++ {
			ticketRules[i] = true
		}
		for i := util.MustAtoi(numRange2[0]); i <= util.MustAtoi(numRange2[1]); i++ {
			ticketRules[i] = true
		}
	}
	for i, line := range input {
		if strings.HasPrefix(line, "your") {
			for _, num := range strings.Split(input[i+1], ",") {
				yourTicket = append(yourTicket, util.MustAtoi(num))
			}
			stopPoint = i + 1
			break
		}
	}
	for i := stopPoint + 1; i < len(input); i++ {
		if strings.HasPrefix(input[i], "nearby") || input[i] == "" {
			continue
		}
		for _, num := range strings.Split(input[i], ",") {
			nearbyTickets = append(nearbyTickets, util.MustAtoi(num))
		}
	}
	for _, num := range nearbyTickets {
		if ok := ticketRules[num]; !ok {
			errors = append(errors, num)
			result += num
		}
	}
	return result, errors
}

func matchMaps(m, n map[string]bool) map[string]bool {
	missingElements := make(map[string]bool)
	for field := range m {
		if ok := n[field]; !ok {
			missingElements[field] = false
		}
	}
	return missingElements
}

func (r rule) checkIfInRange(n int) bool {
	return (n >= r.a && n <= r.b) || (n >= r.c && n <= r.d)
}

func multiplyDepartures(fieldmappings []string, yourTicket []int) int {
	result := 1
	for i, name := range fieldmappings {
		if strings.Fields(name)[0] == "departure" {
			result *= yourTicket[i]
		}
	}
	return result
}

func mapFields(input []string) int {
	var newLine int
	var allFields []string
	var yourTicket []int
	ticketRules := make(map[string]rule)
	mapFields := make(map[string][]bool)

	for i, line := range input {
		if line == "" {
			newLine++
		}
		if line == "" || line == "your ticket:" || line == "nearby tickets:" {
			continue
		}

		switch newLine {
		case 0:
			splitLine := strings.Split(line, ": ")
			allFields = append(allFields, splitLine[0])
			var a, b, c, d int
			fmt.Sscanf(splitLine[1], "%d-%d or %d-%d", &a, &b, &c, &d)
			ticketRules[splitLine[0]] = rule{index: i, name: splitLine[0], a: a, b: b, c: c, d: d}
			mapFields[splitLine[0]] = []bool{}
		case 1:
			for _, num := range strings.Split(line, ",") {
				yourTicket = append(yourTicket, util.MustAtoi(num))
			}
			for field := range mapFields {
				mapFields[field] = make([]bool, len(yourTicket))
				for j := range mapFields[field] {
					mapFields[field][j] = true
				}
			}
		case 2:
			var reverseMapFields [][]string
			var validField bool
			for _, num := range strings.Split(line, ",") {
				validField = false
				var reverseFields []string
				for name, rule := range ticketRules {
					if rule.checkIfInRange(util.MustAtoi(num)) {
						validField = true
					} else {
						reverseFields = append(reverseFields, name)
					}
				}
				reverseMapFields = append(reverseMapFields, reverseFields)
				if !validField {
					break
				}
			}
			if validField {
				for j, reverseFields := range reverseMapFields {
					for _, name := range reverseFields {
						mapFields[name][j] = false
					}
				}
			}
		}
	}

	fieldmappings := make([]string, len(ticketRules))
	for i := 0; i < len(fieldmappings); i++ {
		for field, unknownFields := range mapFields {
			var position []int
			for i, unknown := range unknownFields {
				if unknown && fieldmappings[i] == "" {
					position = append(position, i)
				}
			}
			if len(position) == 1 {
				fieldmappings[position[0]] = field
			}
		}
	}
	return multiplyDepartures(fieldmappings, yourTicket)
}
