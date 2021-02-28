package day09

import (
	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

func findInvalidNumber(lines []int) int {
	for x, y := 0, 25; y < len(lines); x, y = x+1, y+1 {
		sumNumbers := make(map[int]bool)
		for _, num1 := range lines[x:y] {
			for _, num2 := range lines[x:y] {
				num := num1 + num2
				sumNumbers[num] = true
			}
		}
		if _, ok := sumNumbers[lines[y]]; ok != true {
			return lines[y]
		}
	}
	return 0
}

// implementation for part 2 with double loops.
func findEncryptionWeakness(lines []int, invalidNumber int) int {
	for i := 0; i < len(lines); i++ {
		sumNumbers := lines[i]
		for j := i + 1; sumNumbers <= invalidNumber; j++ {
			sumNumbers = sumNumbers + lines[j]
			if sumNumbers == invalidNumber {
				min, max := util.MinMax(lines[i:j])
				return min + max
			}
		}
	}
	return 0
}

// second implementation for part 2 with single loop and if statements.
func findEncryptionWeakness2(lines []int, invalidNumber int) int {
	startIndex := 0
	endIndex := 1
	sum := lines[startIndex]
	for {
		if sum > invalidNumber {
			sum -= lines[startIndex]
			startIndex++
		} else if sum < invalidNumber {
			sum += lines[endIndex]
			endIndex++
		} else {
			min, max := util.MinMax(lines[startIndex:endIndex])
			return min + max
		}
	}
}
