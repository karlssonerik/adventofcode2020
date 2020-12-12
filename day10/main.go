package main

import (
	"fmt"
	"sort"

	"github.com/karlssonerik/adventofcode2020/day10/input"
)

var foundArrs = make(map[string]bool)

func main() {
	in := input.Input
	in = append(in, 0)
	sort.Ints(in)
	in = append(in, in[len(in)-1]+3)
	Find3sN1s(in)
	fmt.Println("start ", in)

	// arrs := FindArrangements(in)
	// fmt.Println("arrangements", arrs+1)
	arrs2 := FindArrangements2(in)
	fmt.Println("arrangements2", arrs2)
}

func FindArrangements2(in []int) int {
	currentIn := make([]int, len(in))
	diffArr := make([]int, len(in))
	copy(currentIn, in)
	possibleArrs := 1
	for i, joltage := range in {
		if i == 0 {
			continue
		}
		diffArr[i] = joltage - currentIn[i-1]
	}

	combo1 := 0
	for i, diff := range diffArr {
		if i == 0 {
			continue
		}
		switch diff {
		case 1:
			combo1++
		case 2:
			possibleArrs *= IntCombo(combo1)
			combo1 = 0

		case 3:
			possibleArrs *= IntCombo(combo1)
			combo1 = 0
		default:
			possibleArrs *= IntCombo(combo1)

		}
	}
	return possibleArrs
}

func IntCombo(combo1s int) (pa int) {
	switch combo1s {
	case 0:
		pa = 1
	case 1:
		pa = 1
	case 2:
		pa = 2
	case 3:
		pa = 4
	case 4:
		pa = 7
	default:
		fmt.Println("oops", combo1s)
		pa = 1
	}
	return pa
}

func FindArrangements(in []int) int {
	arrs := 0
	currentIn := make([]int, len(in))
	copy(currentIn, in)
	for i, joltage := range in {
		if i > (len(currentIn) - 3) {
			continue
		}
		if currentIn[i+2]-joltage <= 3 {
			currentInhead := make([]int, len(currentIn[:i+1]))
			copy(currentInhead, currentIn[:i+1])
			newArr := append(currentInhead, currentIn[i+2:]...)
			if !foundArrs[fmt.Sprint("", newArr)] {
				arrs++
				foundArrs[fmt.Sprint("", newArr)] = true
			} else {
				continue
			}

			arrs += FindArrangements(newArr)
		}
	}
	return arrs

}

func Find3sN1s(in []int) {
	threes := 0
	ones := 0
	for i, joltage := range in {
		if i == 0 {
			continue
		}
		switch joltage - in[i-1] {
		case 1:
			ones++
		case 2:
		case 3:
			threes++
		default:
			panic("wat")
		}
	}
	fmt.Println("1s", ones)
	fmt.Println("3s", threes)
}
