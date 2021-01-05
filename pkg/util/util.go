package util

import (
	"log"
	"strconv"
)

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

// Abs takes an int and returns the absolute value.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// MustAtoi takes a string and returns an int or logs a fatal error if convertion fails.
func MustAtoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not convert number: %v", err)
	}
	return num
}

// MustParseInt takes a string and returns an int or logs a fatal error if convertion fails.
func MustParseInt(s string, base int, bitSize int) int64 {
	num, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		log.Fatalf("could not parse number: %v", err)
	}
	return num
}

// GCD returns the greatest common divisor by the Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// ReverseString takes a string and returns it reversed
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
