package main

import (
	"log"
)

func part2() {
	similarity := 0
	for i := 0; i < len(list1); i++ {
		v1 := list1[i]

		instances := findInstances(v1)

		similarity += v1 * instances
	}

	log.Println(similarity)
}

func findInstances(num int) int {
	instances := 0

	for _, v := range list2 {
		if v == num {
			instances++
		}
	}

	return instances
}
