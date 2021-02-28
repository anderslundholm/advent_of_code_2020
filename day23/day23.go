package day23

import (
	"container/ring"
)

func parseInput(input []string) []int {
	var cups []int
	for _, char := range input[0] {
		cups = append(cups, int(char-'0'))
	}
	return cups

}

func crabCups(input []string, part2 bool) (res int) {
	labels := parseInput(input)
	cupsLen := len(labels)
	moves := 100
	if part2 {
		cupsLen = 1000000
		moves = 10000000
	}

	cups := ring.New(cupsLen)
	cupsMap := map[int]*ring.Ring{}

	for i := 1; i <= cupsLen; i++ {
		if cups.Value = i; i <= len(labels) {
			cups.Value = labels[i-1]
		}
		cupsMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for i := 0; i < moves; i++ {
		selectedCups := cups.Unlink(3)
		destinationCup := (cupsLen+cups.Value.(int)-2)%cupsLen + 1
		removedCups := map[int]bool{}

		for i := 0; i < 3; i++ {
			removedCups[selectedCups.Value.(int)] = true
			selectedCups = selectedCups.Next()
		}
		for removedCups[destinationCup] {
			destinationCup = (cupsLen+destinationCup-2)%cupsLen + 1
		}
		cupsMap[destinationCup].Link(selectedCups)
		cups = cups.Next()
	}

	if cupsLen > len(labels) {
		return cupsMap[1].Next().Value.(int) * cupsMap[1].Move(2).Value.(int)
	}
	cupsMap[1].Unlink(len(labels) - 1).Do(func(p interface{}) {
		res = res*10 + p.(int)
	})
	return
}
