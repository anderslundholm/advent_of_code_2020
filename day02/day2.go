package day02

import (
	"strconv"
	"strings"
)

type dbEntry struct {
	min       int
	max       int
	character string
	password  string
}

func parseFile(line string) (dbEntry, error) {
	var entry dbEntry
	subString := strings.Split(line, " ")

	minMax := strings.Split(subString[0], "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return entry, err
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return entry, err
	}

	entry = dbEntry{
		min:       min,
		max:       max,
		character: strings.TrimRight(subString[1], ":"),
		password:  subString[2],
	}

	return entry, nil
}
