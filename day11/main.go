package main

import (
	"fmt"

	"github.com/karlssonerik/adventofcode2020/day11/input"
)

func main() {
	in := input.Seats
	nextIn := make([][]rune, len(in))
	copy(nextIn, in)
	//printSeats(in)
	iter := 0
	for {
		iter++
		var changed bool
		nextIn, changed = moveSeat(nextIn)
		//	printSeats(nextIn)
		if !changed {
			fmt.Println("taken seats ", countTakenSeats(nextIn))
			fmt.Println("iters", iter)
			break
		}
	}
}

func countTakenSeats(in [][]rune) int {
	count := 0
	for _, row := range in {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

func printSeats(in [][]rune) {
	for _, row := range in {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func moveSeat(in [][]rune) (newIn [][]rune, changed bool) {
	nextIn := make([][]rune, len(in))
	change := false
	for i, row := range in {
		nextIn[i] = make([]rune, len(row))
		for j, cell := range row {
			nextIn[i][j] = evalCell(cell, j, i, in)
			if !change {
				change = (nextIn[i][j] != in[i][j])
			}
		}
	}
	return nextIn, change
}

func evalCell(status rune, x, y int, oldIn [][]rune) (newRune rune) {
	newRune = status
	switch status {
	case '.':
		return '.'
	case '#':
		if getRumpedSeats(x, y, oldIn) >= 5 {
			newRune = 'L'
		}
	case 'L':
		if getRumpedSeats(x, y, oldIn) == 0 {
			newRune = '#'
		}
	}
	return
}

func getRumpedSeat(x, y int, dir dirFunc, oldIn [][]rune) int {
	x, y = dir(x, y)
	seatIsTaken := 0
	if x >= 0 && x < len(oldIn[0]) &&
		y >= 0 && y < len(oldIn) {
		seat := oldIn[y][x]
		switch seat {
		case '#':
			seatIsTaken = 1
		case '.':
			seatIsTaken = getRumpedSeat(x, y, dir, oldIn)
		case 'L':
			seatIsTaken = 0
		}
	}

	return seatIsTaken

}

type dirFunc func(int, int) (int, int)

func goNW(x, y int) (int, int) {
	return x - 1, y - 1
}
func goN(x, y int) (int, int) {
	return x, y - 1
}
func goNE(x, y int) (int, int) {
	return x + 1, y - 1
}

func goW(x, y int) (int, int) {
	return x - 1, y
}
func goE(x, y int) (int, int) {
	return x + 1, y
}

func goSW(x, y int) (int, int) {
	return x - 1, y + 1
}
func goS(x, y int) (int, int) {
	return x, y + 1
}
func goSE(x, y int) (int, int) {
	return x + 1, y + 1
}

func getRumpedSeats(x, y int, oldIn [][]rune) int {
	takenSeats := 0
	takenSeats += getRumpedSeat(x, y, goNW, oldIn)
	takenSeats += getRumpedSeat(x, y, goN, oldIn)
	takenSeats += getRumpedSeat(x, y, goNE, oldIn)
	takenSeats += getRumpedSeat(x, y, goW, oldIn)
	takenSeats += getRumpedSeat(x, y, goE, oldIn)
	takenSeats += getRumpedSeat(x, y, goSW, oldIn)
	takenSeats += getRumpedSeat(x, y, goS, oldIn)
	takenSeats += getRumpedSeat(x, y, goSE, oldIn)

	return takenSeats
}
