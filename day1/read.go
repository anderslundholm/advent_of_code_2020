package day1

import (
	"bufio"
	"io"
	"strconv"
)

func readInts(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)
	var result []int

	for s.Scan() {
		x, err := strconv.Atoi(s.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, s.Err()
}
