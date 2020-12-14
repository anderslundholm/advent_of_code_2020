package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

type bitValue int

var discoveredBitMasks = make(map[string]bool)

func (value *bitValue) applyMask(zeroMask string, oneMask string) {
	zeroBitMask := bitValue(util.MustParseInt(zeroMask, 2, 0))
	oneBitMask := bitValue(util.MustParseInt(oneMask, 2, 0))

	*value |= zeroBitMask
	*value &= oneBitMask
}

func (value *bitValue) applyMaskV2(oneMask string) {
	oneBitMask := bitValue(util.MustParseInt(oneMask, 2, 0))

	*value &= oneBitMask
}

func genearateFluxMasks(buildMask, originalMask, address string) []string {
	if len(originalMask) == 0 {
		return []string{buildMask}
	}
	currentDigit := string(originalMask[0])
	remainingDigits := string(originalMask[1:])
	switch currentDigit {
	case "0":
		buildMask = buildMask + string(address[len(buildMask)])
		return genearateFluxMasks(buildMask, remainingDigits, address)
	case "1":
		return genearateFluxMasks(buildMask+"1", remainingDigits, address)
	case "X":
		xCase := genearateFluxMasks(buildMask+"0", remainingDigits, address)
		oCase := genearateFluxMasks(buildMask+"1", remainingDigits, address)
		return append(xCase, oCase...)
	default:
		panic(fmt.Errorf("unknown character: %s", currentDigit))
	}
}

func dockingEmulatorV2(input []string) int {
	var result int
	memoryMap := make(map[string]bitValue)
	var mask string
	for _, line := range input {
		splitLine := strings.Split(line, " = ")
		if strings.HasPrefix(splitLine[0], "mask") {
			mask = splitLine[1]
		} else if strings.HasPrefix(line, "mem[") {
			memoryaddressess := int64(util.MustAtoi(strings.TrimSuffix(strings.TrimPrefix(splitLine[0], "mem["), "]")))
			value := bitValue(util.MustAtoi(splitLine[1]))
			address := strconv.FormatInt(memoryaddressess, 2)
			address = strings.Repeat("0", 36-len(address)) + address
			generatedMasks := genearateFluxMasks("", mask, address)
			for _, fluxMask := range generatedMasks {
				memoryMap[fluxMask] = value
			}
		}
	}
	for _, v := range memoryMap {
		result += int(v)
	}
	return result
}

func dockingEmulator(input []string) int {
	var result int
	memoryMap := make(map[int]bitValue)
	var zeroMask string
	var oneMask string
	for _, line := range input {
		splitLine := strings.Split(line, " = ")

		if strings.HasPrefix(splitLine[0], "mask") {
			mask := splitLine[1]
			zeroMask = strings.ReplaceAll(mask, "X", "0")
			oneMask = strings.ReplaceAll(mask, "X", "1")
		} else if strings.HasPrefix(line, "mem[") {
			memoryaddressess := util.MustAtoi(strings.TrimSuffix(strings.TrimPrefix(splitLine[0], "mem["), "]"))
			value := strings.TrimSpace(splitLine[1])
			numValue := bitValue(util.MustAtoi(value))
			numValue.applyMask(zeroMask, oneMask)
			memoryMap[memoryaddressess] = numValue
		}
	}
	for _, v := range memoryMap {
		result += int(v)
	}
	return result
}
