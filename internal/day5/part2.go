package day5

import (
	"fmt"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

// https://adventofcode.com/2022/day/5
func SolvePart2(input string) int {
	result := 0
	for line := range stringUtils.GetLine(input) {
		fmt.Printf("line: %s\n", line)
	}
	return result
}
