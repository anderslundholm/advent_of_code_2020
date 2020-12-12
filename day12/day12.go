package day12

import (
	"log"
	"strconv"

	"github.com/anderslundholm/advent_of_code_2020/pkg/util"
)

type direction struct {
	north   string
	east    string
	south   string
	west    string
	left    string
	right   string
	forward string
}

type coordinates struct {
	x int
	y int
}

type location struct {
	north int
	east  int
	south int
	west  int
}

var (
	compass = map[string]int{
		"N": 0,
		"E": 90,
		"S": 180,
		"W": 270,
	}
	nav = direction{
		north:   "N",
		east:    "E",
		south:   "S",
		west:    "W",
		left:    "L",
		right:   "R",
		forward: "F",
	}
)

// Part 1
func changeShipDirection(compassDirection int, directionChange int) string {
	compassDirection = (compassDirection + directionChange) % 360
	for {
		if compassDirection < 0 {
			compassDirection = compassDirection + 360
		} else {
			break
		}
	}
	for k, v := range compass {
		if v == compassDirection {
			return k
		}
	}
	return "Err"
}

// Part 1
func navigateShip(input []string) int {
	var c coordinates
	currentDirection := nav.east
	for _, line := range input {
		dir := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Could not convert to int: %v", err)
		}
		if dir == nav.forward {
			dir = currentDirection
		}
		switch string(dir) {
		case nav.north:
			c.y += num
		case nav.south:
			c.y -= num
		case nav.east:
			c.x += num
		case nav.west:
			c.x -= num
		case nav.left:
			currentDirection = changeShipDirection(compass[currentDirection], -num)
		case nav.right:
			currentDirection = changeShipDirection(compass[currentDirection], num)
		default:
			log.Fatalln("Unknown input.")
		}
	}
	return util.Abs(c.x) + util.Abs(c.y)
}

// Part 2
func changeWaypointDirection(dir string, num int, waypointLocation *map[string]int) {
	turns := num / 90
	for i := 0; i < turns; i++ {
		tmp := (*waypointLocation)[nav.north]
		if dir == nav.right {
			(*waypointLocation)[nav.north] = (*waypointLocation)[nav.west]
			(*waypointLocation)[nav.west] = (*waypointLocation)[nav.south]
			(*waypointLocation)[nav.south] = (*waypointLocation)[nav.east]
			(*waypointLocation)[nav.east] = tmp
		} else if dir == nav.left {
			(*waypointLocation)[nav.north] = (*waypointLocation)[nav.east]
			(*waypointLocation)[nav.east] = (*waypointLocation)[nav.south]
			(*waypointLocation)[nav.south] = (*waypointLocation)[nav.west]
			(*waypointLocation)[nav.west] = tmp
		}
	}
}

// Part 2
func navigateWaypoint(input []string) int {
	shipLocation := map[string]int{"N": 0, "E": 0, "S": 0, "W": 0}
	waypointLocation := map[string]int{"N": 1, "E": 10, "S": 0, "W": 0}
	for _, line := range input {
		dir := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Could not convert to int: %v", err)
		}
		switch dir {
		case nav.forward:
			for k, v := range waypointLocation {
				shipLocation[k] += num * v
			}
		case nav.left:
			changeWaypointDirection(dir, num, &waypointLocation)
		case nav.right:
			changeWaypointDirection(dir, num, &waypointLocation)
		default:
			waypointLocation[dir] += num
		}
	}
	return util.Abs(shipLocation[nav.north]-shipLocation[nav.south]) + util.Abs(shipLocation[nav.east]-shipLocation[nav.west])
}
