package day25

import (
	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

const divider = 20201227

func parseInput(input []string) (int, int) {
	return util.MustAtoi(input[0]), util.MustAtoi(input[1])
}

func getEncryptionKey(input []string) int {
	cardPubKey, doorPubKey := parseInput(input)

	subjectNumber := 7
	loop := 0
	for value := 1; value != cardPubKey; loop++ {
		value = value * subjectNumber % divider
	}

	result := 1
	for i := 0; i < loop; i++ {
		result = result * doorPubKey % divider
	}

	return result

}
