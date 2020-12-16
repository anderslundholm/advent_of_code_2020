package day16

import (
	"fmt"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

var (
	allFields        = make(map[string]bool)
	ticketRules      = make(map[int]map[string]bool)
	reverseMapFields = make(map[int]map[string]bool)
)

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
			// fmt.Println("#######not found", field)

		}
	}
	return missingElements
}

func getErrorsV2(input []string) (int, []int) {
	var result int
	var yourTicket []int
	var stopPoint int
	var errors []int

	var nearbyTickets [][]int

	for i, line := range input {
		if input[i] == "" {
			break
		}
		splitLine := strings.Split(line, ": ")
		allFields[splitLine[0]] = true
		numRange := strings.Split(splitLine[1], " or ")
		numRange1 := strings.Split(numRange[0], "-")
		numRange2 := strings.Split(numRange[1], "-")
		ticketRules[i] = map[string]bool{}
		for j := util.MustAtoi(numRange1[0]); j <= util.MustAtoi(numRange1[1]); j++ {
			ticketRules[j] = map[string]bool{splitLine[0]: true}

		}
		for j := util.MustAtoi(numRange2[0]); j <= util.MustAtoi(numRange2[1]); j++ {
			ticketRules[j] = map[string]bool{splitLine[0]: true}
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
		var tickets []int
		for _, s := range strings.Split(input[i], ",") {
			num := util.MustAtoi(s)
			if _, ok := ticketRules[num]; !ok {
				errors = append(errors, num)
				result += num
			} else {
				tickets = append(tickets, num)
			}
		}
		nearbyTickets = append(nearbyTickets, tickets)
	}

	for i := range nearbyTickets {
		for _, j := range nearbyTickets[i] {
			missingElements := matchMaps(allFields, ticketRules[j])
			reverseMapFields[i] = missingElements
		}
	}

	mapFields(reverseMapFields, allFields)

	// fmt.Println("reverse", reverseMapFields)
	return result, errors
}

func mapFields(reverseMapFields map[int]map[string]bool, allFields map[string]bool) {
	asdf := make(map[int]string)
	for len(asdf) < len(allFields) {
		for i := range reverseMapFields {
			for j := range allFields {
				if ok := reverseMapFields[i][j]; ok {
					fmt.Println("found something")
				}
			}
		}
	}
}
