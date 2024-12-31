package main

import (
	"log"
	"math"
)

func part1() {
	sum := 0
	for i := 0; i < len(list1); i++ {
		distance := int(math.Abs(float64(list1[i] - list2[i])))

		sum += distance
	}

	log.Println("Sum of total distances:", sum)
}
