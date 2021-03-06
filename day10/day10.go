package day10

import (
	"sort"
)

var discoveredArrangements = make(map[int]int)

func prepList(lines *[]int) {
	sort.Ints(*lines)
	*lines = append(*lines, (*lines)[len(*lines)-1]+3)
	*lines = append([]int{0}, *lines...)
}

func findJoltageChain(lines []int) int {
	oneJoltDiff := 0
	threeJoltDiff := 0
	prepList(&lines)
	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1]-lines[i] == 1 {
			oneJoltDiff++
		} else if lines[i+1]-lines[i] == 3 {
			threeJoltDiff++
		}
	}

	result := oneJoltDiff * threeJoltDiff
	return result
}

func countChargerArrangements(lines []int) int {
	prepList(&lines)
	result := recursiveCount(0, lines)
	return result
}

// implementation for part 2 with recursive function.
func recursiveCount(recurseIndex int, lines []int) int {
	if recurseIndex == len(lines)-1 {
		return 1
	} else if _, ok := discoveredArrangements[recurseIndex]; ok {
		return discoveredArrangements[recurseIndex]
	}

	result := 0
	for i := recurseIndex + 1; i < len(lines); i++ {
		if lines[i]-lines[recurseIndex] <= 3 {
			result += recursiveCount(i, lines)
		}
	}
	discoveredArrangements[recurseIndex] = result
	return result
}

// second implementation for part 2 with single loop and if statements.
func checkAdjacentNumbers(lines []int) int {
	prepList(&lines)
	subsum := 1
	count := 0
	for i := range lines {
		if i == len(lines)-1 {
			break
		}
		diff := lines[i+1] - lines[i]
		if diff == 3 && count > 1 {
			if count == 2 {
				subsum = subsum * 2
			} else if count == 3 {
				subsum = subsum * 4
			} else if count == 4 {
				subsum = subsum * 7
			} else {
				subsum = subsum * count
			}
			count = 0
		}
		if diff == 3 && count <= 1 {
			count = 0
		} else if diff == 1 {
			count++
		}
	}

	return subsum
}
