package main

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

var list []string

func main() {
	list = loadInput()

	part1()
	part2()
}

func loadInput() []string {
	var text string
	if len(os.Args) > 1 && os.Args[1] == "--sample" {
		text = sample
	} else {
		text = input
	}

	return strings.Split(text, "\n")
}

func isCharacter(character byte, x, y int) bool {
	// check for out of bounds
	if y > len(list) || y < 0 {
		return false
	}
	if x > len(list[y])-1 || x < 0 {
		return false
	}
	return list[y][x] == character
}
