package day11

import (
	"reflect"
)

const (
	floor        = "."
	emptySeat    = "L"
	occupiedSeat = "#"
)

type coord struct {
	x int
	y int
}

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

func countNewAdjacentOccupied(input [][]string, i int, j int) int {
	adjacentOccupied := 0
	var possibleCoordinates []coord
	possibleCoordinates = []coord{
		coord{x: 1, y: 0}, coord{x: 1, y: 1}, coord{x: 0, y: 1}, coord{x: -1, y: 1},
		coord{x: -1, y: 0}, coord{x: -1, y: -1}, coord{x: 0, y: -1}, coord{x: 1, y: -1}}
	for _, coordinate := range possibleCoordinates {
		cc := coord{x: i + coordinate.x, y: j + coordinate.y}
		discoveredSeat := false
		for !discoveredSeat && (0 <= cc.x && cc.x < len(input)) && (0 <= cc.y && cc.y < len(input[0])) {
			seat := input[cc.x][cc.y]
			switch seat {
			case occupiedSeat:
				adjacentOccupied++
				discoveredSeat = true
			case emptySeat:
				discoveredSeat = true
			case floor:
				cc.x += coordinate.x
				cc.y += coordinate.y
			}

		}
	}

	return adjacentOccupied
}

func newModelProcess(input *[][]string) int {
	for {
		var seatsToOccupy []coord
		var seatsToEmpty []coord
		for i := range *input {
			for j := range (*input)[0] {
				seat := (*input)[i][j]
				if seat == floor {
					continue
				}
				adjacentOccupied := countNewAdjacentOccupied(*input, i, j)
				if seat == emptySeat && adjacentOccupied == 0 {
					c := coord{x: i, y: j}
					seatsToOccupy = append(seatsToOccupy, c)
				} else if seat == occupiedSeat && adjacentOccupied >= 5 {
					c := coord{x: i, y: j}
					seatsToEmpty = append(seatsToEmpty, c)
				}
			}
		}
		for _, seat := range seatsToOccupy {
			(*input)[seat.x][seat.y] = occupiedSeat
		}
		for _, seat := range seatsToEmpty {
			(*input)[seat.x][seat.y] = emptySeat
		}

		if len(seatsToEmpty)+len(seatsToOccupy) == 0 {
			var result int
			for i := range *input {
				for _, seat := range (*input)[i] {
					if seat == occupiedSeat {
						result++
					}
				}
			}
			return result
		}
	}
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
