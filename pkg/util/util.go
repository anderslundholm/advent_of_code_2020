package util

// MinMax takes an array of numbers and returns the minimum and the maximum numbers.
func MinMax(numbersArray []int) (int, int) {
	max := numbersArray[0]
	min := numbersArray[0]
	for _, value := range numbersArray {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}

	}
	return min, max
}
