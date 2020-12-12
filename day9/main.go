package main

import (
	"fmt"
	"sort"

	"github.com/karlssonerik/adventofcode2020/day9/input"
)

func main() {
	fmt.Println("asdasdas")
	in := input.Input
	preambleLen := 25
	hackerNumber := 0
	for i := preambleLen; i < len(in); i++ {
		preamble := in[i-preambleLen : i]
		if !isValid(in[i], preamble) {
			fmt.Println("not valid", in[i], preamble)
			hackerNumber = in[i]
			break
		}
	}
	for i := 0; i < len(in); i++ {
		checkForHackerSum(hackerNumber, i, in)
	}
}

func checkForHackerSum(hackerNumber, index int, in []int) {
	currentSet := []int{}
	sum := 0
	for _, aNum := range in[index:] {
		currentSet = append(currentSet, aNum)
		sum += aNum
		if sum == hackerNumber && len(currentSet) > 1 {
			sort.Ints(currentSet)
			fmt.Println("Sum found!", currentSet)
			fmt.Println("hacker SUM", currentSet[0]+currentSet[len(currentSet)-1])
		}
		if sum > hackerNumber {
			break
		}
	}

}

func isValid(anum int, preamble []int) bool {
	for idx, preamb := range preamble {
		for j := idx + 1; j < len(preamble); j++ {
			if preamble[j]+preamb == anum {
				fmt.Printf("%d is equal to %d + %d\n", anum, preamb, preamble[j])
				return true
			}
		}
	}
	fmt.Println("end of loop")
	return false

}
