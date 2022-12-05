package day3

import (
	"fmt"
	"strings"
)

// https://adventofcode.com/2022/day/3
// Given an input of the ruck scack, compute the priorities
func SolvePart2(input string) int {
	totalScore := 0
	lines := strings.Split(input, "\n")

	for group := range getGroups(lines) {
		badge, prirority := getGroupPriority(group)
		totalScore += prirority
		fmt.Printf("On groups: %v, Found: %d (%s)\n", group, prirority, badge)
	}

	return totalScore
}

func getGroupPriority(group [3]string) (string, int) {
	// Get the common characters
	common := getCommonInLines(group)
	// Get the priority of the common character
	priority := getPriority(common)
	return common, priority
}

func getCommonInLines(group [3]string) string {
	cache := make(map[string]bool)
	for _, character := range group[0] {
		cache[string(character)] = true
	}

	cache2 := make(map[string]bool)
	for _, character := range group[1] {
		if _, ok := cache[string(character)]; ok {
			cache2[string(character)] = true
		}
	}

	for _, character := range group[2] {
		if _, ok := cache2[string(character)]; ok {
			return string(character)
		}
	}

	return ""
}

func getGroups(line []string) chan [3]string {
	channel := make(chan [3]string)

	go func() {
		l := len(line)
		for i := 0; i < l; i += 3 {
			channel <- [3]string{line[i], line[i+1], line[i+2]}
		}
		close(channel)
	}()

	return channel
}
