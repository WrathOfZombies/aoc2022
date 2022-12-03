package day1

import (
	"math"
	"strings"

	"github.com/wrathofzombies/aoc2022/internal/common/stringUtils"
)

// https://adventofcode.com/2022/day/1
// Given an input of the calories carried by each elf, finds the elf carrying the most
// amount of calories and returns the sum of the calories being carried by that elf
func SolvePart1(input string) int {
	line := strings.Split(input, "\n")
	globalMax := 0
	localMax := 0

	for _, item := range line {
		// If the line is blank, then we move on to the next elf
		if item == "" {
			// Compare the globalMax with our localMax
			// Update the globalMax to the localMax, if the latter is greater
			globalMax = int(math.Max(float64(globalMax), float64(localMax)))
			localMax = 0
			continue
		}

		// Otherwise parse the line to get the calorie count
		// and update our localmax
		calorie := stringUtils.ParseInt(item)
		localMax += calorie
	}

	// flush
	globalMax = int(math.Max(float64(globalMax), float64(localMax)))
	localMax = 0

	return globalMax
}
