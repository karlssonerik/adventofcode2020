package main

import "fmt"

func main() {
	bigCount := 0
	bigCountB := 0
	for _, group := range input {
		uniqueAnswers := make(map[rune]bool)
		answersCount := make(map[rune]int)
		counter := 0
		members := 1

		for _, r := range group {
			if r == 10 {
				members++
				continue
			}
			if !uniqueAnswers[r] {
				uniqueAnswers[r] = true
			}
			answersCount[r] = answersCount[r] + 1
		}
		for range uniqueAnswers {
			counter++
		}
		bigCount += counter

		for _, v := range answersCount {
			if v == members {
				bigCountB++
			}
		}

	}
	fmt.Println("Part1", bigCount)
	fmt.Println("Part2", bigCountB)
}

var example = []string{
	`abc`,

	`a
b
c`,

	`ab
ac`,

	`a
a
a
a`,

	`b`,
}
