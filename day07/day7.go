package day07

import (
	"log"
	"strconv"
	"strings"
)

var myBag = bagType{
	description: "shiny",
	color:       "gold",
}

type bagType struct {
	description string
	color       string
}

type bagRules map[bagType]map[bagType]int

func (rules bagRules) checkIfBagContainsBag(root bagType) bool {
	if _, ok := rules[root][myBag]; ok {
		return true
	}
	// Recusively check all bags
	for bag := range rules[root] {
		if rules.checkIfBagContainsBag(bag) {
			return true
		}
	}

	return false
}

func (rules bagRules) countBags(bag bagType) int {
	count := 1
	for innerBag, numBags := range rules[bag] {
		count += numBags * rules.countBags(innerBag)
	}

	return count
}

func parseBags(lines []string) bagRules {
	allBags := bagRules{}
	for _, line := range lines {
		splitLine := strings.Split(line, " bags contain ")

		rawRootBag := splitLine[0]
		innerBags := splitLine[1]
		splitRawRootBag := strings.Split(rawRootBag, " ")
		rootBag := bagType{
			description: splitRawRootBag[0],
			color:       splitRawRootBag[1],
		}

		allBags[rootBag] = make(map[bagType]int)

		if strings.HasSuffix(line, "no other bags.") != true {
			innerBags := strings.TrimSuffix(innerBags, " bags contain no other bags.")
			for _, rawInnerBag := range strings.Split(innerBags, ", ") {
				splitRawInnerBag := strings.Split(rawInnerBag, " ")
				num, err := strconv.Atoi(splitRawInnerBag[0])
				if err != nil {
					log.Fatalf("Could not convert to int: %v", err)
				}
				innerBag := bagType{
					description: splitRawInnerBag[1],
					color:       splitRawInnerBag[2],
				}
				allBags[rootBag][innerBag] = num

			}
		}
	}

	return allBags
}
