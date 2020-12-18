package main

import (
	"fmt"
	"strconv"
	"strings"
)

//
//Example1: 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
//Example2: 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
//Example3: ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
var Example = []string{
	"5 + " + doMath("8 * 3 + 9 + 3 * 4 * 3") + "",
	"5 * 9 * " + doMath("7 * 3 * 3 + 9 * 3 + "+doMath("8 + 6 * 4")+"") + "",
	"" + doMath(""+doMath("2 + 4 * 9")+" * "+doMath("6 + 9 * 8 + 6")+" + 6") + " + 2 + 4 * 2",
}

/*
1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
2 * 3 + (4 * 5) becomes 46.
5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
*/
var Example2 = []string{
	"1 + " + doMath("2 * 3") + " + " + doMath("4 * "+doMath("5 + 6")+"") + "",
	"2 * 3 + " + doMath("4 * 5") + "",
	"5 + " + doMath("8 * 3 + 9 + 3 * 4 * 3") + "",
	"5 * 9 * " + doMath("7 * 3 * 3 + 9 * 3 + "+doMath("8 + 6 * 4")+"") + "",
	"" + doMath(""+doMath("2 + 4 * 9")+" * "+doMath("6 + 9 * 8 + 6")+" + 6") + " + 2 + 4 * 2",
}

func main() {
	in := Input
	sum := 0
	for _, s := range in {
		dm := doMath(s)
		current, err := strconv.Atoi(dm)
		if err != nil {
			panic("NaN")
		}
		//fmt.Println("Math1!", current)
		sum += current

	}
	fmt.Println("Math!", sum)
}

type mathFunc func(int, int) int

func multi(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func doMath(s string) string {
	if len(strings.TrimSpace(s)) == 0 {
		return ""
	}

	countmulti := strings.Count(s, "*")
	sssss := ""
	if countmulti > 0 {
		ssss := strings.Split(s, "*")
		for idx, r := range ssss {
			r = strings.TrimSpace(r)
			partSum := doMath_part1(r, false)
			if strings.TrimSpace(partSum) == "" {
				continue
			}
			if idx == 0 {
				sssss = partSum
			} else {
				sssss += " * " + partSum
			}
		}
	} else {
		sssss = s
	}
	sssss = strings.TrimSpace(sssss)
	return doMath_part1(sssss, false)
}

func doMath_part1(s string, print bool) string {
	if print {
		fmt.Println("doing ", s)
	}
	if len(strings.TrimSpace(s)) == 0 {
		return ""
	}
	ss := ""
	firstNumFound := false
	i := 0
	var currentFunc mathFunc
	sss := strings.Split(s, " ")
	for _, r := range sss {
		if r == " " || r == "" {
			continue
		} else if r == "*" {
			currentFunc = multi
		} else if r == "+" {
			currentFunc = add
		} else {
			num, err := strconv.Atoi(r)
			if err != nil {
				panic("NAN")
			}
			if !firstNumFound {
				firstNumFound = true
				i = num
			} else {
				i = currentFunc(i, num)
			}
		}

	}

	ss = fmt.Sprintf("%d", i)
	return ss
}
