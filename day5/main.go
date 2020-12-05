package main

import "fmt"

func main() {
	colMax := 128
	rowsMax := 8
	maxSeatID := 0
	seats := [128][8]bool{}
	for _, sequence := range input {
		rArr := []rune(sequence)
		cols := iArr(colMax)
		rows := iArr(rowsMax)
		for _, letter := range rArr {
			rHalf := len(rows) / 2
			cHalf := len(cols) / 2
			switch letter {
			case 'B':
				cols = cols[cHalf:]
			case 'F':
				cols = cols[:cHalf]
			case 'L':
				rows = rows[:rHalf]
			case 'R':
				rows = rows[rHalf:]
			default:
				panic("lol wut")
			}
		}
		seatID := (cols[0] * 8) + rows[0]
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		seats[cols[0]][rows[0]] = true
		fmt.Println("seat: ", cols[0], rows[0], seatID)
	}
	fmt.Println("max ", maxSeatID)
	for i, seatrow := range seats {
		for j, taken := range seatrow {
			if !taken && i > 8 && i < 117 {
				fmt.Println("Free seat!", i, j, ((i * 8) + j))
			}
		}
	}

}

func iArr(max int) []int {
	iArr := make([]int, max)
	for i := 0; i < max; i++ {
		iArr[i] = i
	}
	return iArr
}

func seats() [128][8]bool {
	return [128][8]bool{}
}
