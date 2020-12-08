package day8

import (
	"log"
	"strconv"
	"strings"
)

const (
	nopInstruction = "nop"
	accInstruction = "acc"
	jmpInstruction = "jmp"
)

// Finds loop in the program and returns the acumilated "acc" value.
func findLoop(lines []string) int {
	accValue := 0
	index := 0
	visitedLines := make(map[int]bool)
	for {
		if _, ok := visitedLines[index]; ok {
			return accValue
		}
		splitLines := strings.Split(lines[index], " ")
		instruction := splitLines[0]
		value, err := strconv.Atoi(splitLines[1])
		if err != nil {
			log.Fatalf("Cannot convert value: %v\n", err)
		}

		if instruction == accInstruction {
			accValue += value
			visitedLines[index] = true
			index++
		} else if instruction == jmpInstruction {
			visitedLines[index] = true
			index += value
		} else {
			visitedLines[index] = true
			index++
		}

	}

}

//
// Test to make recursive solution
//

// func fixLoop(lines []string, index int, accValue int, changed bool, visitedLines map[int]bool, changedInstructions map[int]bool) (result int, breakRecursion bool) {

// 	if breakRecursion {
// 		return accValue, breakRecursion
// 	}
// 	if index >= 8 {
// 		fmt.Println("helloooooooooo")
// 		breakRecursion = true
// 		return accValue, breakRecursion
// 	}
// 	if _, ok := visitedLines[index]; ok {
// 		visitedLines := make(map[int]bool)
// 		accValue, breakRecursion = fixLoop(lines, 0, 0, false, visitedLines, changedInstructions)
// 	}
// 	splitLines := strings.Split(lines[index], " ")
// 	instruction := splitLines[0]
// 	value, err := strconv.Atoi(splitLines[1])
// 	if err != nil {
// 		log.Fatalf("Cannot convert value: %v\n", err)
// 	}

// 	if instruction == accInstruction {
// 		accValue += value
// 		visitedLines[index] = true
// 		fmt.Println(index, instruction, value, changed, visitedLines, changedInstructions)
// 		accValue, breakRecursion = fixLoop(lines, index+1, accValue, changed, visitedLines, changedInstructions)
// 	} else if instruction == jmpInstruction {
// 		if _, ok := changedInstructions[index]; ok != true && changed != true {
// 			changedInstructions[index] = true
// 			changed = true
// 			visitedLines[index] = true

// 			accValue, breakRecursion = fixLoop(lines, index+1, 0, changed, visitedLines, changedInstructions)
// 		} else {
// 			visitedLines[index] = true
// 			accValue, breakRecursion = fixLoop(lines, index+value, accValue, changed, visitedLines, changedInstructions)
// 		}
// 	} else if instruction == nopInstruction {
// 		if _, ok := changedInstructions[index]; ok != true && value != 0 && changed != true {
// 			changedInstructions[index] = true
// 			changed = true
// 			visitedLines[index] = true
// 			accValue, breakRecursion = fixLoop(lines, index+value, 0, changed, visitedLines, changedInstructions)
// 		} else {
// 			visitedLines[index] = true
// 			accValue, breakRecursion = fixLoop(lines, index+1, accValue, changed, visitedLines, changedInstructions)
// 		}
// 	}
// 	return accValue, breakRecursion
// }

func runProgram2(lines []string) int {
	var changedInstructions = make(map[int]bool)
	accValue := 0
	index := 0
	visitedLines := make(map[int]bool)
	var changed bool
	for {
		if index >= len(lines) {
			return accValue
		}
		if _, ok := visitedLines[index]; ok {
			accValue = 0
			index = 0
			visitedLines = make(map[int]bool)
			changed = false
			continue
		}
		splitLines := strings.Split(lines[index], " ")
		instruction := splitLines[0]
		value, err := strconv.Atoi(splitLines[1])
		if err != nil {
			log.Fatalf("Cannot convert value: %v\n", err)
		}

		if instruction == accInstruction {
			accValue += value
			visitedLines[index] = true
			index++
		} else if instruction == jmpInstruction {
			if _, ok := changedInstructions[index]; ok != true && value != 0 && changed != true {
				changedInstructions[index] = true
				changed = true
				visitedLines[index] = true
				index++
			} else {
				visitedLines[index] = true
				index += value
			}
		} else if instruction == nopInstruction {
			if _, ok := changedInstructions[index]; ok != true && value != 0 && changed != true {
				changedInstructions[index] = true
				changed = true
				visitedLines[index] = true
				index += value
			} else {
				visitedLines[index] = true
				index++
			}
		}
	}
}
