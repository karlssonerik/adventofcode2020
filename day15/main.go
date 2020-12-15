package main

import "fmt"

func main() {
	memory := map[int]int{
		2:  0,
		0:  1,
		6:  2,
		12: 3,
		1:  4,
		3:  5,
	}
	// example := map[int]int{0: 0, 3: 1, 6: 2}
	// memory = example
	last := 0
	for i := len(memory); i < 30000000; i++ {
		idx, ok := memory[last]
		memory[last] = i
		if i == (30000000 - 1) {
			fmt.Println("last-", last)
		}
		if !ok {
			last = 0
		} else {
			last = i - idx
		}

	}
}
