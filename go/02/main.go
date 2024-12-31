package main

import (
	_ "embed"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

var list [][]int

func main() {
	var err error
	list, err = loadInput()
	if err != nil {
		log.Fatal(err)
	}

	part1()
	part2()
}

func loadInput() ([][]int, error) {
	var inputTrimmed string
	var output [][]int
	if len(os.Args) > 1 && os.Args[1] == "--sample" {
		inputTrimmed = strings.Trim(sample, "\n")
	} else {
		inputTrimmed = strings.Trim(input, "\n")
	}

	lines := strings.Split(inputTrimmed, "\n")

	for _, line := range lines {
		inputs := strings.Split(line, " ")

		list := []int{}
		for i, v := range inputs {
			converted, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Error parsing input at line: " + strconv.Itoa(i))
			}

			list = append(list, converted)
		}
		output = append(output, list)
	}

	return output, nil
}
