package day15

// findInSlice takes a slice of ints and an int and returns the differens between
// the index and the int in the list if it exists, otherwise it returns 0.
func findInSlice(list []int, i int) int {
	t := len(list)
	for j, n := range list[:len(list)-1] {
		if list[i] == n {
			t = j
		}
	}
	if t < len(list) {
		return i - t
	}
	return 0
}

func getLastNumber(list []int, turns int) int {
	memory := make([]int, 0, turns)
	memory = append(memory, list...)
	for i := 0; i < turns-1; i++ {
		if i == len(memory)-1 {
			memory = append(memory, findInSlice(memory, i))
		}
	}
	return memory[len(memory)-1]
}

func getLastNumberV2(list []int, turns int) int {
	spokenNumbers := make(map[int]int)
	var lastNumber int
	for i, num := range list {
		lastNumber = num
		spokenNumbers[lastNumber] = i + 1
	}

	for i := len(spokenNumbers); i <= turns-1; i++ {
		if j, ok := spokenNumbers[lastNumber]; ok {
			spokenNumbers[lastNumber] = i
			lastNumber = i - j
		} else {
			spokenNumbers[lastNumber] = i
			lastNumber = 0
		}
	}
	return lastNumber
}
