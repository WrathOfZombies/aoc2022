package day6_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day6"
)

// Tests the example input
func TestExamplePart2(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s", tc.input)
		t.Run(testname, func(t *testing.T) {
			if result := SolvePart2(tc.input); result != tc.expected {
				t.Errorf("Testing part 1: Output %d not equal to expected %d", result, tc.expected)
			}
		})
	}
}

// Tests the input file
func TestInputPart2(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED int = 2260
	if result := SolvePart2(string(input)); !reflect.DeepEqual(result, EXPECTED) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", result, EXPECTED)
	}
}
