package day17

var (
	cubes = make(map[coord]bool)
)

type coord struct {
	x, y, z, w int
}

func (c coord) findNeighbours() []coord {
	var neighbours []coord
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if c.w == -1 {
					if x == 0 && y == 0 && z == 0 {
						continue
					}
					neighbours = append(neighbours, coord{x: c.x + x, y: c.y + y, z: c.z + z})
				} else {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						neighbours = append(neighbours, coord{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w})
					}
				}
			}
		}
	}
	return neighbours
}

func parseInput(input []string, cycles, middleCoord int) {
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				cubes[coord{x: x + cycles, y: y + cycles, z: middleCoord}] = true
			}
		}
	}
}

func parse4DInput(input []string, cycles, middleCoord int) {
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				cubes[coord{x: x + cycles, y: y + cycles, z: middleCoord, w: middleCoord}] = true
			}
		}
	}
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
			count++
		}
	}
	return count
}

func (c coord) runCycle(active bool) bool {
	neighbours := c.findNeighbours()

	activeNeighbours := countActiveNeighbours(neighbours)
	if active && (activeNeighbours == 2 || activeNeighbours == 3) {
		return true
	} else if !active && activeNeighbours == 3 {
		return true
	}
	return false
}

func initGrid(xSize, ySize, cycles int) int {
	xSize = xSize + cycles*2
	ySize = ySize + cycles*2
	zSize := 1 + cycles*2
	start := 1 + (zSize / 2)
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			for z := 0; z < zSize; z++ {
				cubes[coord{x: x, y: y, z: z}] = false
			}
		}
	}
	return start
}

func init4DGrid(xSize, ySize, cycles int) int {
	zwSize := 1 + cycles*2
	middleCoord := zwSize / 2
	for x := 0; x < xSize+cycles*2; x++ {
		for y := 0; y < ySize+cycles*2; y++ {
			for z := 0; z < zwSize; z++ {
				for w := 0; w < zwSize; w++ {
					cubes[coord{x: x, y: y, z: z, w: w}] = false
				}
			}
		}
	}
	return middleCoord
}

func simulateCycles(input []string, cycles int) int {
	start := initGrid(len(input[0]), len(input), cycles)
	parseInput(input, cycles, start)
	newCubes := make(map[coord]bool)
	for i := 0; i < cycles; i++ {
		for coordinates, active := range cubes {
			newCubes[coordinates] = coordinates.runCycle(active)
		}
		for k, v := range newCubes {
			cubes[k] = v
		}
	}
	return countAllActive()
}

func simulate4DCycles(input []string, cycles int) int {
	middleCoord := init4DGrid(len(input[0]), len(input), cycles)
	parse4DInput(input, cycles, middleCoord)
	newCubes := make(map[coord]bool)
	for i := 0; i < cycles; i++ {
		for coordinates, active := range cubes {
			newCubes[coordinates] = coordinates.runCycle(active)
		}
		for k, v := range newCubes {
			cubes[k] = v
		}
	}
	return countAllActive()
}
