package main

import (
	"fmt"
	"strings"
)

type field struct {
	name       string
	intervalls []intervall
}
type intervall struct {
	min int
	max int
}

func main() {
	sum := 0
	validTickets := [][]int{}
	fieldIndexes := make(map[string][]int)
	for _, tick := range Input_nearby_tickets {
		currentTickectfieldIndexes := make(map[string][]int)
		isTicketValid := true
		for fieldIdx, v := range tick {
			vChecked := false
			for _, f := range InputFields {
				for _, intervall := range f.intervalls {
					if v >= intervall.min && v <= intervall.max {
						vChecked = true
						indexes, ok := currentTickectfieldIndexes[f.name]
						if !ok {
							indexes = []int{fieldIdx}
						} else {
							indexes = appendIfunique(indexes, fieldIdx)
						}

						currentTickectfieldIndexes[f.name] = indexes
					}
				}

			}
			if !vChecked {
				isTicketValid = false
				sum += v
			}
		}

		if isTicketValid {
			for k, v := range currentTickectfieldIndexes {
				indexes, ok := fieldIndexes[k]
				if !ok {
					indexes = v
				} else {
					indexes = appendIfUniques(indexes, v)
				}
				fieldIndexes[k] = indexes
			}
			validTickets = append(validTickets, tick)
		}
	}
	fmt.Println("loopy sum", sum)
	prunedIndexes := make(map[string][]int)
	for k, v := range fieldIndexes {
		newV := []int{}
		for _, index := range v {
			add := true
			for _, tick := range validTickets {
				intLen := len(InputFields2[k])
				failingInter := 0
				for _, intervall := range InputFields2[k] {
					if !(tick[index] >= intervall.min && tick[index] <= intervall.max) {
						failingInter++
					}
				}
				if failingInter == intLen {
					add = false
					break
				}
			}
			if add {
				newV = append(newV, index)
			} else {

			}
		}
		prunedIndexes[k] = newV
	}

	FinalFieldNIndex := make(map[string]int)
	FoundIndex := make(map[int]bool)
	done := false
	fieldsByIndexesCopy := make(map[string][]int)
	for !done {
		for k, v := range prunedIndexes {
			newV := []int{}
			for _, i := range v {
				if !FoundIndex[i] {

					newV = append(newV, i)
				}
			}
			v = newV
			if len(v) == 1 {
				//fmt.Println("found", k, v[0])
				FinalFieldNIndex[k] = v[0]
				FoundIndex[v[0]] = true
			}
			fieldsByIndexesCopy[k] = v
		}
		prunedIndexes = fieldsByIndexesCopy
		done = len(FinalFieldNIndex) == len(InputFields)
	}
	fmt.Println("done", FinalFieldNIndex)
	endSum := 1
	for k, v := range FinalFieldNIndex {
		if strings.HasPrefix(k, "departure") {
			endSum *= Input_your_ticket[v]
		}

	}
	fmt.Printf("end %d\n", endSum)
}

func appendIfunique(slice []int, v int) (res []int) {
	res = make([]int, len(slice))
	copy(res, slice)
	add := true
	for _, vInSlice := range slice {
		if v == vInSlice {
			add = false
			break
		}
	}
	if add {
		res = append(res, v)
	}
	return res
}

func appendIfUniques(slice []int, vs []int) []int {
	res := make([]int, len(slice))
	copy(res, slice)
	for _, v := range vs {
		res = appendIfunique(res, v)
	}
	return res
}
