package day11

import (
	"reflect"
)

const (
	floor        = "."
	emptySeat    = "L"
	occupiedSeat = "#"
)

func getRowNum(index int, rowLength int) int {
	rowIndex := index % rowLength
	rowNum := (index - rowIndex) / rowLength
	return rowNum
}

func countAdjacentOccupied(input map[int]string, i int, rowLength int) int {
	adjacentOccupied := 0
	var adjacentSeats []int
	rowIndex := i % rowLength

	// Not top
	if i >= rowLength {
		adjacentSeats = append(adjacentSeats, i-rowLength)
		// Not Left
		if rowIndex != 0 {
			adjacentSeats = append(adjacentSeats, i-rowLength-1)
		}
		// Not Right
		if rowIndex != rowLength-1 {
			adjacentSeats = append(adjacentSeats, i-rowLength+1)
		}
	}
	// Not bottom
	if i < len(input)-rowLength {
		adjacentSeats = append(adjacentSeats, i+rowLength)

		// Not Left
		if rowIndex != 0 {
			adjacentSeats = append(adjacentSeats, i+rowLength-1)
		}
		// Not Right
		if rowIndex != rowLength-1 {
			adjacentSeats = append(adjacentSeats, i+rowLength+1)
		}
	}
	// Not Left
	if rowIndex != 0 {
		adjacentSeats = append(adjacentSeats, i-1)
	}
	// Not Right
	if rowIndex != rowLength-1 {
		adjacentSeats = append(adjacentSeats, i+1)
	}

	for _, x := range adjacentSeats {
		if input[x] == occupiedSeat {
			adjacentOccupied++
		}
	}

	return adjacentOccupied
}

func modelProcess(input *map[int]string, rowLength int) int {
	var countOccupiedSeats int
	seatMap := make(map[int]string)
	for {

		countOccupiedSeats = 0
		for i := range *input {

			adjacentOccupied := countAdjacentOccupied(*input, i, rowLength)
			switch (*input)[i] {
			case emptySeat:
				if adjacentOccupied == 0 {
					seatMap[i] = occupiedSeat
					countOccupiedSeats++
				} else {
					seatMap[i] = emptySeat
				}
			case occupiedSeat:
				if adjacentOccupied >= 4 {
					seatMap[i] = emptySeat
					countOccupiedSeats--
				} else {
					seatMap[i] = occupiedSeat
					countOccupiedSeats++
				}
			case floor:
				seatMap[i] = floor
			}
		}

		if reflect.DeepEqual(*input, seatMap) {
			break
		}

		for k, v := range seatMap {
			(*input)[k] = v
		}
	}
	return countOccupiedSeats
}
