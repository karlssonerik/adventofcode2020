package main

import (
	"errors"
	"fmt"

	"github.com/karlssonerik/adventofcode2020/day8/input"
)

func main() {
	instructions := input.Input

	iteration := 0
	for {
		acc, err := runBootUp(instructions, iteration)
		if err == nil {
			fmt.Println("accumolator ", acc, iteration)
			break
		}
		iteration++
	}

}

func runBootUp(instructions []input.Instruction, iterationNo int) (int, error) {
	index := 0
	accumolator := 0
	executedIns := make(map[int]bool)
	changer := 0
	for {
		if executedIns[index] {
			return 0, errors.New("Loop")
		}
		if index >= len(instructions) {
			return accumolator, nil
		}
		executedIns[index] = true
		currrentIns := instructions[index]
		switch currrentIns.Op {
		case "nop":
			if changer == iterationNo {
				index += currrentIns.Arg
			} else {
				index++
			}
			changer++
		case "acc":
			accumolator += currrentIns.Arg
			index++
		case "jmp":
			if changer == iterationNo {
				index++
			} else {
				index += currrentIns.Arg
			}
			changer++
		}
	}
}
