package day20

import (
	"strings"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

var (
	imageTiles []tile
)

type tile struct {
	id      int
	tile    []string
	edges   []string
	flipped bool
}

func multiplyCorners(matchedTiles map[int][]string) int {
	result := 1
	for id, edges := range matchedTiles {
		unique := make(map[string]bool)
		for _, s := range edges {
			unique[s] = true
		}
		if len(unique) == 2 {
			result *= id
		}
	}
	return result
}

func parseTiles(input []string) {
	edge := 0
	var t tile
	for _, line := range input {
		if strings.HasPrefix(line, "Tile ") {
			t = tile{}
			t.id = util.MustAtoi(strings.TrimSuffix(strings.TrimPrefix(line, "Tile "), ":"))
			edge = 1
			continue
		}
		if line == "" {
			edge = 1
			continue
		}
		switch edge {
		case 1:
			t.edges = append(t.edges, string(line[0]))
			t.edges = append(t.edges, string(line[9]))
			t.edges = append(t.edges, line)
			edge++
			continue
		case 10:
			t.edges = append(t.edges, line)
			imageTiles = append(imageTiles, t)
		}
		t.edges[0] = t.edges[0] + string(line[0])
		t.edges[1] = t.edges[1] + string(line[9])
		edge++
	}
}

func assembleImage(input []string) int {
	matchedTiles := make(map[int][]string)
	parseTiles(input)

	for _, tile := range imageTiles {
		for _, v := range tile.edges {
			for _, tile2 := range imageTiles {
				if tile.id == tile2.id {
					continue
				}
				for _, v2 := range tile2.edges {
					flippedEdge := util.ReverseString(v2)
					if v == v2 || v == flippedEdge {
						matchedTiles[tile.id] = append(matchedTiles[tile.id], v2)
						matchedTiles[tile2.id] = append(matchedTiles[tile2.id], v)
					}
				}
			}
		}
	}
	return multiplyCorners(matchedTiles)
}
