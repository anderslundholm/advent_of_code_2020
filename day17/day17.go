package day17

import "fmt"

var (
	cubes = make(map[coord]bool)
)

type coord struct {
	x, y, z int
}

func (c coord) findNeighbours() []coord {
	var neighbours []coord
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighbours = append(neighbours, coord{
					x: c.x + x,
					y: c.y + y,
					z: c.z + z,
				})
			}
		}
	}
	// fmt.Println(neighbours)
	return neighbours
}

func parseInput(input []string, startZ int) {
	// cube := make(map[coord]bool)
	for y, line := range input {
		for x, char := range line {
			switch char {
			case '#':
				cubes[coord{x: x, y: y, z: startZ}] = true
			}
		}
	}
	// fmt.Println(cubes)
	// return cube
}

func countAllActive() int {
	count := 0
	for _, v := range cubes {
		if v {
			count++
		}
	}
	return count
}

func countActiveNeighbours(neighbours []coord) int {
	count := 0
	for _, neighbour := range neighbours {

		if cubes[neighbour] == true {
			// fmt.Println(neighbour, cubes[neighbour])
			count++
		}
	}
	return count
}

func (c coord) runCycle(active bool) bool {
	neighbours := c.findNeighbours()
	activeNeighbours := countActiveNeighbours(neighbours)
	if active && (activeNeighbours == 2 || activeNeighbours == 3) {
		// fmt.Println("1")
		return true
	} else if !active && activeNeighbours == 3 {
		// fmt.Println("2")
		return true
	}
	return false

}

func initGrid(xSize, ySize, zSize int) {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			for z := 0; z < zSize; z++ {
				cubes[coord{x: x, y: y, z: z}] = false
			}
		}
	}
}

func simulateCycles(input []string, cycles int) int {
	xSize := len(input[0]) + cycles*2
	ySize := len(input) + cycles*2
	zSize := 1 + cycles*2
	startZ := 1 + (zSize / 2)
	fmt.Println(startZ, zSize)
	initGrid(xSize, ySize, zSize)
	parseInput(input, startZ)

	newCubes := make(map[coord]bool)
	// fmt.Println(cubes)

	for i := 0; i < cycles; i++ {
		// newCube := make(map[coord]bool)
		for coordinates, active := range cubes {
			// fmt.Println(cubes)

			newCubes[coordinates] = coordinates.runCycle(active)
			if newCubes[coordinates] {
				// fmt.Println(coordinates.x, coordinates.y, coordinates.z)
			}
		}
		for k, v := range newCubes {
			cubes[k] = v
		}
		fmt.Println("#############################################################")
	}
	// fmt.Println(cubes)
	return countAllActive()
}
