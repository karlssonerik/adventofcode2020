package main

import (
	"fmt"
)

var arrivalAtBusStop = 1000677
var a2 = 939
var b2 = []int{7, 13, -1, -1, 59, -1, 31, 19}

var busIds = []int{29, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 41, -1, -1, -1, -1, -1, -1, -1, -1, -1, 661, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 13, 17, -1, -1, -1, -1, -1, -1, -1, -1, 23, -1, -1, -1, -1, -1, -1, -1, 521, -1, -1, -1, -1, -1, 37, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 19}

var busindex = make(map[int]int)
var waitTimes = make(map[int]int)

func main() {

	in := arrivalAtBusStop
	busses := busIds
	lowestwaittime := 100000000000
	nearestbus := 0
	for _, bus := range busses {
		if bus == -1 {
			continue
		}
		wt := (bus - (in % bus))
		//	fmt.Println(" ", bus, (in % bus), wt)
		if wt < lowestwaittime {
			lowestwaittime = wt
			nearestbus = bus
		}
	}
	fmt.Println("next bus at ", lowestwaittime*nearestbus)
	mahIndex := 0
	for idx, bus := range busses {
		if bus == -1 {
			continue
		}
		busindex[bus] = idx
		waitTimes[bus] = (bus - (in % bus))
		mahIndex++
	}
	lowest := int64(0)
	for idx, bus := range busses {
		if bus == -1 {
			continue
		}
		if idx != 0 {
			lowest = commonNumbers(busses, lowest, int64(idx))
		}

	}

	fmt.Println("LowestCommon", lowest)

}

func commonNumbers(busses []int, previousLowest, endIdx int64) int64 {
	lowest := int64(0)
	bigNumber := int64(1)
	for i := int64(0); i <= endIdx; i++ {
		if busses[i] == -1 {
			continue
		}
		bigNumber *= int64(busses[i])
	}

	if previousLowest == 0 {
		for i := int64(0); i < bigNumber; i++ {
			bigB := int64(busses[endIdx])
			if (i%bigB == (bigB - endIdx)) && i%int64(busses[0]) == 0 {
				lowest = i
				break
			}
		}
	} else {

		for i := previousLowest; i < bigNumber; i += (bigNumber / int64(busses[endIdx])) {
			busTimestampAfterInital := (int64(busses[endIdx]) - endIdx) % int64(busses[endIdx])
			if busTimestampAfterInital < 0 {
				busTimestampAfterInital += int64(busses[endIdx])
			}
			if i%int64(busses[endIdx]) == busTimestampAfterInital {
				lowest = i
			}
		}

	}
	return lowest
}
