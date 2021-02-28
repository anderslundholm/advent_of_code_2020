package day24

import (
	"regexp"
)

var (
	coordinates = map[string]coord{
		"e": {2, 0}, "w": {-2, 0},
		"ne": {1, -1}, "nw": {-1, -1},
		"se": {1, 1}, "sw": {-1, 1},
	}
)

type coord struct {
	x, y int
}

func (c coord) add(c2 coord) coord {
	return coord{c.x + c2.x, c.y + c2.y}
}

func flipTiles(input []string, part2 bool) int {
	blackTiles := map[coord]struct{}{}
	for _, line := range input {
		c := coord{0, 0}
		for _, s := range regexp.MustCompile(`(e|se|sw|w|nw|ne)`).FindAllString(line, -1) {
			c = c.add(coordinates[s])
		}

		if _, ok := blackTiles[c]; ok {
			delete(blackTiles, c)
		} else {
			blackTiles[c] = struct{}{}
		}
	}
	if !part2 {
		return len(blackTiles)
	}
	return livingArtExhibit(blackTiles)
}

func livingArtExhibit(blackTiles map[coord]struct{}) int {
	for i := 0; i < 100; i++ {
		neighbour := map[coord]int{}
		for coordinate := range blackTiles {
			for _, c := range coordinates {
				neighbour[coordinate.add(c)]++
			}
		}

		flip := map[coord]struct{}{}
		for coordinate, n := range neighbour {
			if _, ok := blackTiles[coordinate]; ok && n == 1 || n == 2 {
				flip[coordinate] = struct{}{}
			}
		}
		blackTiles = flip
	}
	return len(blackTiles)
}
