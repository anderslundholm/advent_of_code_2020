package reader

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// OpenFile ---
func openFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open file: %v\n", err)
	}
	return f
}

// ReadInts reads and returns int from file.
func ReadInts(filePath string) ([]int, error) {
	f := openFile(filePath)
	defer f.Close()

	s := bufio.NewScanner(f)
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

// ReadLines reads from file line by line and returns a slice of strings.
func ReadLines(filePath string) ([]string, error) {
	f := openFile(filePath)
	defer f.Close()
	s := bufio.NewScanner(f)
	var result []string

	for s.Scan() {
		result = append(result, s.Text())
	}
	return result, s.Err()
}
