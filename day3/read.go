package day3

import (
	"bufio"
	"io"
)

func readLines(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var result []string

	for s.Scan() {
		result = append(result, s.Text())
	}
	return result, s.Err()
}
