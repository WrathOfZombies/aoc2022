package day4

import (
	"fmt"
	"strings"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

// https://adventofcode.com/2022/day/4
// Given an input find the number of ranges with any overlaps
func SolvePart2(input string) int {
	totalScore := 0
	for line := range stringUtils.GetLine(input) {
		pair := strings.Split(line, ",")
		l, r := pair[0], pair[1]
		match := checkAnyOverlap(l, r)
		if match {
			totalScore++
			fmt.Printf("Matched line: %s, totalScore: %d\n", line, totalScore)
		}
	}
	return totalScore
}

func checkAnyOverlap(l, r string) bool {
	lStart, lEnd := getRange(l)
	rStart, rEnd := getRange(r)

	fmt.Printf("lStart: %d, lEnd: %d, rStart: %d, rEnd: %d\n", lStart, lEnd, rStart, rEnd)

	// If left is completely outside right
	if lEnd < rStart {
		fmt.Printf("Left is completely outside right\n")
		return false
	}

	// If right is completely outside left
	if rEnd < lStart {
		fmt.Printf("Right is completely outside left\n")
		return false
	}

	return true
}
