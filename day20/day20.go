package day20

import (
	"fmt"
	"image"
	"math"
	"regexp"
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

// Tile ...
type Tile []string

func (t Tile) col(i int) (col string) {
	for _, s := range t {
		col += string(s[i])
	}
	return
}

func (t Tile) edges() []string {
	return []string{t[0], t.col(len(t[0]) - 1), t[len(t)-1], t.col(0)}
}

func (t Tile) rotations() []Tile {
	rot := make([]Tile, 8)
	for r := 0; r < 8; r += 2 {
		for _, s := range t {
			rot[r] = append(rot[r], util.ReverseString(s))
		}
		for i := range t {
			rot[r+1] = append(rot[r+1], util.ReverseString(t.col(i)))
		}
		t = rot[r+1]
	}
	return rot
}

func part2(input string) int {

	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	tiles := map[int]Tile{}
	counts := map[string]int{}
	for _, s := range split {
		var id int
		fmt.Sscanf(s, "Tile %d:", &id)
		tiles[id] = strings.Split(s, "\n")[1:]
		for _, e := range tiles[id].edges() {
			counts[e]++
			counts[util.ReverseString(e)]++
		}
	}

	imageSize, tileSize := int(math.Sqrt(float64(len(tiles)))), 10
	img := make(Tile, imageSize*(tileSize-2))
	order := map[image.Point]Tile{}

	part1 := 1
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
		findTile:
			for i, t := range tiles {
				for _, r := range t.rotations() {
					if (y == 0 && counts[r[0]] == 1 ||
						y != 0 && r[0] == order[image.Point{x, y - 1}][tileSize-1]) &&
						(x == 0 && counts[r.col(0)] == 1 ||
							x != 0 && r.col(0) == order[image.Point{x - 1, y}].col(tileSize-1)) {
						if (y == 0 || y == imageSize-1) && (x == 0 || x == imageSize-1) {
							part1 *= i
						}

						for i := 0; i < tileSize-2; i++ {
							img[(tileSize-2)*y+i] += r[i+1][1 : tileSize-1]
						}

						order[image.Point{x, y}] = r
						delete(tiles, i)
						break findTile
					}
				}
			}
		}
	}
	fmt.Println(part1)

	mon := []string{"..................#.", "#....##....##....###", ".#..#..#..#..#..#..."}
	nmon := 0
	for _, r := range img.rotations() {
		for y := 0; y < len(r)-len(mon); y++ {
		findMon:
			for x := 0; x < len(r[0])-len(mon[0]); x++ {
				for i, s := range mon {
					if match, _ := regexp.MatchString(s, r[y+i][x:x+len(s)]); !match {
						continue findMon
					}
				}
				nmon++
			}
		}
	}
	return strings.Count(strings.Join(img, ""), "#") - nmon*strings.Count(strings.Join(mon, ""), "#")
}
