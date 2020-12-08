package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2020/pkg/timer"
)

func runProgram1(lines []string) int {

	accValue := 0
	currrentLine := 0
	visitedLines := make(map[int]bool)
	for {
		if _, ok := visitedLines[currrentLine]; ok {
			return accValue
		}
		splitLines := strings.Split(lines[currrentLine], " ")
		instruction := splitLines[0]
		value, err := strconv.Atoi(splitLines[1])
		if err != nil {
			log.Fatalf("Cannot convert value: %v\n", err)
		}

		if instruction == accInstruction {
			accValue += value
			visitedLines[currrentLine] = true
			currrentLine++
		} else if instruction == jmpInstruction {
			visitedLines[currrentLine] = true
			currrentLine += value
		} else {
			visitedLines[currrentLine] = true
			currrentLine++
		}

	}

}

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day8/input.txt")
	if err != nil {
		log.Fatalf("Could not read lines: %v\n", err)
	}

	var allLines []string
	for _, line := range lines {
		allLines = append(allLines, line)

	}
	result := findLoop(allLines)

	fmt.Printf("Day8 Part1 result: %v\n", result)
}
