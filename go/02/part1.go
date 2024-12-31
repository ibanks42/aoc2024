package main

import (
	"log"
	"math"
	"slices"
)

func part1() {
	safeReports := 0
	for _, v := range list {
		if isSafe(v) {
			safeReports++
		}
	}
	log.Println("Safe reports part1:", safeReports)
}

func part2() {
	safeReports := 0
	for _, v := range list {
		if isSafePart2(v) {
			safeReports++
		}
	}
	log.Println("Safe reports part2:", safeReports)
}

func isSafe(nums []int) bool {
	increasing := nums[0]-nums[1] > 0

	for i := 0; i < len(nums)-1; i++ {
		if !isNextSafe(increasing, nums[i], nums[i+1]) {
			return false
		}
	}
	return true
}

func isSafePart2(nums []int) bool {
	hasRemovedLevel := false
	increasing := nums[0]-nums[1] > 0

	for i := 0; i < len(nums)-1; i++ {
		safe := isNextSafe(increasing, nums[i], nums[i+1])
		if !safe && !hasRemovedLevel {
			nums = slices.Delete(nums, i, i+1)
			if i+1 > len(nums)-1 {
				return true
			}

			hasRemovedLevel = true
			if !isNextSafe(increasing, nums[i], nums[i+1]) {
				return false
			}
		}
	}

	return true
}

func isNextSafe(increasing bool, one int, two int) bool {
	val := one - two
	if val < 0 && increasing {
		return false
	}
	if val > 0 && !increasing {
		return false
	}

	abs := math.Abs(float64(val))
	if abs == 0 || abs > 3 {
		return false
	}
	return true
}
