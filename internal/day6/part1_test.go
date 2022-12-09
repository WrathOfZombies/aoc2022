package day6_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day6"
)

// Tests the example input
func TestExamplePart1(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s", tc.input)
		t.Run(testname, func(t *testing.T) {
			if result := SolvePart1(tc.input); result != tc.expected {
				t.Errorf("Testing part 1: Output %d not equal to expected %d", result, tc.expected)
			}
		})
	}
}

// Tests the input file
func TestInputPart1(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED int = 0
	if result := SolvePart1(string(input)); !reflect.DeepEqual(result, EXPECTED) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", result, EXPECTED)
	}
}
