package day3

import (
	"fmt"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

// https://adventofcode.com/2022/day/3
// Given an input of the ruck scack, compute the priorities
func SolvePart1(input string) int {
	totalScore := 0

	// Split the input into lines
	for line := range stringUtils.GetLine(input) {
		common := getCommon(line)
		priority := getPriority(common)
		totalScore += priority
		fmt.Printf("On line: %s, Found: %d (%s)\n", line, priority, common)
	}

	return totalScore
}

// Given a map, update the count of the key or add it if it doesn't exist
func addOrUpdate(m map[string]int, k string) map[string]int {
	if _, ok := m[k]; ok {
		m[k]++
	} else {
		m[k] = 1
	}
	return m
}

// Given a line, return a common of the left and right characters; split down the middle
func getCommon(line string) string {
	trackLeft := make(map[string]int)
	trackRight := make(map[string]int)

	mid := len(line) / 2

	for i := 0; i < mid; i++ {
		left := line[i]
		right := line[i+mid]
		trackLeft = addOrUpdate(trackLeft, string(left))
		trackRight = addOrUpdate(trackRight, string(right))

		if _, found := trackLeft[string(right)]; found {
			return string(right)
		}

		if _, found := trackRight[string(left)]; found {
			return string(left)
		}
	}

	return ""
}

// Given a string, return the priority of the character
func getPriority(s string) int {
	ascii := int(s[0])

	if s >= "a" && s <= "z" {
		return (ascii % 97) + 1
	}

	if s >= "A" && s <= "Z" {
		return (ascii % 65) + 27
	}

	return 0
}
