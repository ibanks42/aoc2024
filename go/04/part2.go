package main

import "log"

func part2() {
	xmasCount := 0

	for y := 0; y < len(list); y++ {
		for x := 0; x < len(list[y]); x++ {
			char := list[y][x]

			if char == 'A' && isMas(x, y) {
				xmasCount++
			}
		}
	}

	log.Println("XMAS count", xmasCount)
}

func isMas(x, y int) bool {
	isMas := false
	isMas2 := false

	// check 1 & 2 first
	// 1 . .
	// . . .
	// . . 2
	if isCharacter('M', x-1, y-1) && isCharacter('S', x+1, y+1) {
		isMas = true
	} else if isCharacter('S', x-1, y-1) && isCharacter('M', x+1, y+1) {
		isMas = true
	}

	// if first check isn't MAS or SAM then we can quit and return false
	if !isMas {
		return false
	}

	// check 3 & 4
	// . . 3
	// . . .
	// 4 . .
	if isCharacter('M', x+1, y-1) && isCharacter('S', x-1, y+1) {
		isMas2 = true
	} else if isCharacter('S', x+1, y-1) && isCharacter('M', x-1, y+1) {
		isMas2 = true
	}

	// return both are MAS or SAM
	return isMas && isMas2
}
