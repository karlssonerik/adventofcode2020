package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlssonerik/adventofcode2020/day7/input"
)

var rules = make(map[string]map[string]int)

func main() {

	for _, rawrule := range input.Input {
		//fmt.Printf("doing %s«\n", rawrule)
		rawrule = strings.Trim(rawrule, ".")
		s := strings.Split(rawrule, "contain")
		rule := make(map[string]int)
		ss := strings.Split(s[1], ",")
		mainBag := strings.TrimSuffix(s[0], " bags ")
		for _, rulePart := range ss {
			rulePart = strings.TrimSpace(rulePart)
			rulePart = strings.TrimSuffix(rulePart, " bags")
			rulePart = strings.TrimSuffix(rulePart, " bag")

			if rulePart == "no other" {
				continue
			}
			ruleParts := strings.SplitN(rulePart, " ", 2)
			units, err := strconv.Atoi(ruleParts[0])
			if err != nil {
				panic(err)
			}
			bag := strings.TrimSuffix(ruleParts[1], " bags")
			bag = strings.TrimSuffix(bag, " bag")
			//fmt.Printf("adding %s:%s:%d«\n", mainBag, bag, units)
			rule[bag] = units

		}
		//fmt.Printf("adding %s:%v«\n", mainBag, rule)
		rules[mainBag] = rule
	}
	counter := 0

	mm := 0
	for k, v := range rules {
		mm++
		if k == "shiny gold" {
			continue
		}

		if checkForGold(v) {
			counter++
		}
	}

	fmt.Println("available bags", counter, mm)

	fmt.Println("count bags", countBags(rules["shiny gold"]))

}

func countBags(bags map[string]int) int {
	counter := 0
	for bag, count := range bags {
		counter += count
		for i := 0; i < count; i++ {
			counter += countBags(rules[bag])
		}
	}
	return counter
}

func checkForGold(bags map[string]int) bool {
	foundgold := false
	for bag := range bags {
		if bag == "shiny gold" {
			foundgold = true
		}
	}
	for bag := range bags {
		if !foundgold {
			foundgold = checkForGold(rules[bag])
		}
	}

	return foundgold
}
