package main

import (
	"log"
)

func part1() {
	xmasCount := 0

	for y := 0; y < len(list); y++ {
		for x := 0; x < len(list[y]); x++ {
			char := list[y][x]
			if char == 'X' {
				xmasCount += checkXmasLine(x, y)
			}
		}
	}

	log.Println("XMAS count", xmasCount)
}

func checkXmasLine(x, y int) int {
	xmasCount := 0

	// check all directions
	// XMAS can be at up, down, left, right, or either diagonal
	// increment xmasCount for each XMAS found
	// but also only check for specific direction in while we're checking
	if checkLineInDirection("up", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("down", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("left", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("right", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("upLeft", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("upRight", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("downLeft", 'X', x, y) {
		xmasCount++
	}
	if checkLineInDirection("downRight", 'X', x, y) {
		xmasCount++
	}

	return xmasCount
}

func checkLineInDirection(direction string, character byte, x, y int) bool {
	// make sure we set the next character to check if we find it
	var nextCharacter byte
	switch character {
	case 'X':
		nextCharacter = 'M'
	case 'M':
		nextCharacter = 'A'
	case 'A':
		nextCharacter = 'S'
	}

	xCheck := x
	yCheck := y

	switch direction {
	case "up":
		yCheck = y - 1
	case "down":
		yCheck = y + 1
	case "left":
		xCheck = x - 1
	case "right":
		xCheck = x + 1
	case "upLeft":
		xCheck = x - 1
		yCheck = y - 1
	case "upRight":
		xCheck = x + 1
		yCheck = y - 1
	case "downLeft":
		xCheck = x - 1
		yCheck = y + 1
	case "downRight":
		xCheck = x + 1
		yCheck = y + 1
	default:
		return false
	}

	if isCharacter(nextCharacter, xCheck, yCheck) {
		// if nextCharacter == correct one and nextCharacter = 'S' then we've found an XMAS.
		if nextCharacter == 'S' {
			return true
		}
		// continue checking the next character in the matrix
		return checkLineInDirection(direction, nextCharacter, xCheck, yCheck)
	}

	return false
}
