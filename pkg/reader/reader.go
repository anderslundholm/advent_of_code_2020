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

// ReadChars reads from file line by line and returns a map of strings.
func ReadChars(filePath string) (map[int]string, int, error) {
	f := openFile(filePath)
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	result := make(map[int]string)
	rowLength := 0
	index := 0
	for s.Scan() {
		if rowLength == 0 && s.Text() == "\n" {
			rowLength = index
		} else if s.Text() != "\n" {
			result[index] = s.Text()
			index++
		}
	}
	return result, rowLength, s.Err()
}

// ReadMultiRunes reads from file line by line and returns a 2d slice of runes.
func ReadMultiLineRunes(filePath string) ([][]string, error) {
	f := openFile(filePath)
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	var result [][]string
	var inner []string
	for s.Scan() {
		if s.Text() == "\n" {
			result = append(result, inner)
			inner = []string{}
		} else {
			inner = append(inner, s.Text())
		}
	}
	result = append(result, inner)
	return result, s.Err()
}
