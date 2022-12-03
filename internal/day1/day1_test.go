package day1_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day1"
)

// Tests the example input and evaluates the Part1 and Part2
func TestExample(t *testing.T) {
	const INPUT string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	const EXPECTED_PART_1 int = 24000
	if resultPart1 := SolvePart1(INPUT); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", resultPart1, EXPECTED_PART_1)
	}

	const EXPECTED_PART_2 int = 45000
	if resultPart2 := SolvePart2(INPUT); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %d not equal to expected %d", resultPart2, EXPECTED_PART_2)
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED_PART_1 int = 75501
	if resultPart1 := SolvePart1(string(input)); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", resultPart1, EXPECTED_PART_1)
	}

	const EXPECTED_PART_2 int = 215594
	if resultPart2 := SolvePart2(string(input)); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %d not equal to expected %d", resultPart2, EXPECTED_PART_2)
	}
}
