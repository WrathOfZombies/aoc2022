package day4

import (
	"fmt"
	"strings"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

// https://adventofcode.com/2022/day/4
// Given an input of the ruck scack, compute the priorities
func SolvePart1(input string) int {
	totalScore := 0
	for line := range stringUtils.GetLine(input) {
		pair := strings.Split(line, ",")
		l, r := pair[0], pair[1]
		match := checkFullOverlap(l, r)
		if match {
			totalScore++
			fmt.Printf("Matched line: %s, totalScore: %d\n", line, totalScore)
		}
	}
	return totalScore
}

func checkFullOverlap(l, r string) bool {
	lStart, lEnd := getRange(l)
	rStart, rEnd := getRange(r)

	// If left overlaps right
	if lStart <= rStart && lEnd >= rEnd {
		return true
	}

	// If right overlaps left
	if rStart <= lStart && rEnd >= lEnd {
		return true
	}

	return false
}

func getRange(input string) (int, int) {
	pair := strings.Split(input, "-")
	return stringUtils.ParseInt(pair[0]), stringUtils.ParseInt(pair[1])
}
