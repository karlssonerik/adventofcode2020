package main

import "fmt"

/*
#......#
##.#..#.
#.#.###.
.##.....
.##.#...
##.#....
#####.#.
##.#.###
*/
type cube struct {
	x int
	y int
	z int
}

var activeCubes = []cube{
	{0, 7, 0}, {7, 7, 0},
	{0, 6, 0}, {1, 6, 0}, {3, 6, 0}, {6, 6, 0},
	{0, 5, 0}, {2, 5, 0}, {4, 5, 0}, {5, 5, 0}, {6, 5, 0},
	{1, 4, 0}, {2, 4, 0},
	{1, 3, 0}, {2, 3, 0}, {4, 3, 0},
	{0, 2, 0}, {1, 2, 0}, {3, 2, 0},
	{0, 1, 0}, {1, 1, 0}, {2, 1, 0}, {3, 1, 0}, {4, 1, 0}, {6, 1, 0},
	{0, 0, 0}, {1, 0, 0}, {3, 0, 0}, {5, 0, 0}, {6, 0, 0}, {7, 0, 0},
}

//var activeCubes = []cube{{0, 0, 0}, {1, 0, 0}, {2, 0, 0}, {2, 1, 0}, {1, 2, 0}}

var activeCubeMap = make(map[cube]bool)

var currentState [][][]rune

func main() {
	noCycle := 6
	for _, aCube := range activeCubes {
		activeCubeMap[aCube] = true
	}
	for i := 0; i < noCycle; i++ {
		pCubs := make(map[cube]bool)
		nextActiveCubes := make(map[cube]bool)
		for _, aCube := range activeCubes {
			pCubs[aCube] = true
			for _, nearCube := range getNeighbours(aCube) {
				pCubs[nearCube] = true
			}
		}

		nextActiveList := []cube{}
		for pCube := range pCubs {
			actives := 0
			for _, nearCube := range getNeighbours(pCube) {
				if activeCubeMap[nearCube] {
					actives++
				}
			}
			if activeCubeMap[pCube] {
				if actives == 2 || actives == 3 {
					nextActiveCubes[pCube] = true
					nextActiveList = append(nextActiveList, pCube)
				}
			} else {
				if actives == 3 {
					nextActiveCubes[pCube] = true
					nextActiveList = append(nextActiveList, pCube)
				}
			}
		}

		activeCubeMap = nextActiveCubes
		activeCubes = nextActiveList
		fmt.Println("asdasd", len(activeCubeMap))
	}

}

func getNeighbours(c cube) []cube {
	return []cube{
		{c.x - 1, c.y, c.z},
		{c.x - 1, c.y - 1, c.z},
		{c.x - 1, c.y + 1, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y + 1, c.z},
		{c.x + 1, c.y, c.z},
		{c.x + 1, c.y - 1, c.z},
		{c.x + 1, c.y + 1, c.z},
		{c.x, c.y, c.z + 1},
		{c.x - 1, c.y, c.z + 1},
		{c.x - 1, c.y - 1, c.z + 1},
		{c.x - 1, c.y + 1, c.z + 1},
		{c.x, c.y - 1, c.z + 1},
		{c.x, c.y + 1, c.z + 1},
		{c.x + 1, c.y, c.z + 1},
		{c.x + 1, c.y - 1, c.z + 1},
		{c.x + 1, c.y + 1, c.z + 1},
		{c.x, c.y, c.z - 1},
		{c.x - 1, c.y, c.z - 1},
		{c.x - 1, c.y - 1, c.z - 1},
		{c.x - 1, c.y + 1, c.z - 1},
		{c.x, c.y - 1, c.z - 1},
		{c.x, c.y + 1, c.z - 1},
		{c.x + 1, c.y, c.z - 1},
		{c.x + 1, c.y - 1, c.z - 1},
		{c.x + 1, c.y + 1, c.z - 1},
	}
}
