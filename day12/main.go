package main

import (
	"fmt"
	"math"
)

type ship struct {
	facing string
	x      int
	y      int
}

type waypoint struct {
	x int
	y int
}

type inst struct {
	op    string
	value int
}

var example = []inst{
	{"F", 10},
	{"N", 3},
	{"F", 7},
	{"R", 90},
	{"F", 11},
}

func main() {
	ship := ship{
		"E", 0, 0,
	}
	waypoint := waypoint{
		10, 1,
	}
	for _, ins := range Input {
		switch ins.op {
		case "F":
			ship.goF(&waypoint, ins.value)
		case "N":
			waypoint.goDir(ins.op, ins.value)
		case "W":
			waypoint.goDir(ins.op, ins.value)
		case "E":
			waypoint.goDir(ins.op, ins.value)
		case "S":
			waypoint.goDir(ins.op, ins.value)
		case "L":
			waypoint.turn(ship, ins.op, ins.value)
		case "R":
			waypoint.turn(ship, ins.op, ins.value)
		default:
			panic("bad op")
		}
	}
	distance := math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))
	fmt.Println("dis ", distance)
}

var dirs = []string{"N", "E", "S", "W"}

func (s *ship) turn(dir string, angle int) {
	dirIndex := 0

	switch s.facing {
	case "N":
		dirIndex = 0
	case "E":
		dirIndex = 1
	case "S":
		dirIndex = 2
	case "W":
		dirIndex = 3
	}

	switch dir {
	case "R":
		times := angle / 90
		dirIndex = (dirIndex + times) % 4
		s.facing = dirs[dirIndex]
	case "L":
		times := angle / 90
		dirIndex = ((dirIndex - times) % 4)
		if dirIndex < 0 {
			dirIndex += 4
		}
		s.facing = dirs[dirIndex]
	}
}

func (s *waypoint) turn(ship ship, dir string, angle int) {
	xF := (s.x - ship.x)
	yF := (s.y - ship.y)
	var newX, newY int
	switch dir {
	case "R":
		switch angle {
		case 90:
			newX = yF
			newY = xF * -1
		case 180:
			newX = xF * -1
			newY = yF * -1
		case 270:
			newX = yF * -1
			newY = xF
		}
	case "L":
		switch angle {
		case 90:
			newX = yF * -1
			newY = xF
		case 180:
			newX = xF * -1
			newY = yF * -1
		case 270:
			newX = yF
			newY = xF * -1
		}
	}
	s.x = int(newX) + ship.x
	s.y = int(newY) + ship.y
}

func (s *waypoint) goDir(dir string, distance int) {
	switch dir {
	case "N":
		s.y += distance
	case "W":
		s.x -= distance
	case "E":
		s.x += distance
	case "S":
		s.y -= distance
	}
}

func (s *ship) goDir(dir string, distance int) {
	switch dir {
	case "N":
		s.y += distance
	case "W":
		s.x -= distance
	case "E":
		s.x += distance
	case "S":
		s.y -= distance
	}
}

func (s *ship) goF(w *waypoint, distance int) {
	//s.goDir(s.facing, distance)
	diffX := w.x - s.x
	diffY := w.y - s.y
	for i := 0; i < distance; i++ {
		s.x += diffX
		s.y += diffY
	}
	w.x = s.x + diffX
	w.y = s.y + diffY

}
