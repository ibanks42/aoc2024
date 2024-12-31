package main

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

var (
	list1 []int
	list2 []int
)

func main() {
	log.Println(os.Args)

	var err error
	list1, list2, err = getSortedInput()
	if err != nil {
		log.Fatal(err)
	}

	part1()
	part2()
}

func getSortedInput() ([]int, []int, error) {
	var inputTrimmed string
	if len(os.Args) > 1 && os.Args[1] == "--sample" {
		inputTrimmed = strings.Trim(sample, "\n")
	} else {
		inputTrimmed = strings.Trim(input, "\n")
	}

	lines := strings.Split(inputTrimmed, "\n")
	list1 := []int{}
	list2 := []int{}

	for i, line := range lines {
		nums := strings.Split(line, "   ")
		str1 := nums[0]
		str2 := nums[1]

		num1, err := strconv.Atoi(str1)
		if err != nil {
			log.Println()
			return nil, nil, errors.New("Error parsing input at line: " + strconv.Itoa(i))
		}
		num2, err := strconv.Atoi(str2)
		if err != nil {
			log.Println("Error parsing input at line: ", i)
			return nil, nil, errors.New("Error parsing input at line: " + strconv.Itoa(i))
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2, nil
}
