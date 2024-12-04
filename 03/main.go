package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed sample2.txt
var sample2 string

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

func part1() {
	sampleOutput := addAll(findAllMultiplications(sample))
	realOutput := addAll(findAllMultiplications(input))

	log.Println("Part1 sample", sampleOutput)
	log.Println("Part1 real", realOutput)
}

func part2() {
	sampleOutput := addAll(findAllMultiplications(removeDonts(sample2)))
	realOutput := addAll(findAllMultiplications(removeDonts(input)))

	log.Println("Part2 sample", sampleOutput)
	log.Println("Part2 real", realOutput)
}

func findAllMultiplications(input string) []int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllStringIndex(input, -1)

	values := []int{}
	for _, v := range matches {
		mul := input[v[0]:v[1]]
		mul = strings.TrimLeft(mul, "mul(")
		mul = strings.TrimRight(mul, ")")
		nums := strings.Split(mul, ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal("Failed to parse input at index " + strconv.Itoa(v[0]))
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal("Failed to parse input at index " + strconv.Itoa(v[0]))
		}

		values = append(values, num1*num2)
	}
	return values
}

func removeDonts(text string) string {
	encapsulatedRegex := regexp.MustCompile(`don't\(\)[\S\s]*?do\(\)`)

	// remove all encapsulated don't/dos
	text = encapsulatedRegex.ReplaceAllLiteralString(text, "")

	// remove last don't if it exists
	standaloneRegex := regexp.MustCompile(`don't\(\)(\S\s]*?)`)
	text = standaloneRegex.ReplaceAllLiteralString(text, "")

	return text
}

func addAll(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
