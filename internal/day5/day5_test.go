package day5_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day5"
)

// Tests the example input and evaluates the Part1 and Part2
func TestExamplePart1(t *testing.T) {
	const INPUT string = ``
	const EXPECTED_PART_1 int = 0
	if resultPart1 := SolvePart1(INPUT); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", resultPart1, EXPECTED_PART_1)
	}
}

func TestExamplePart2(t *testing.T) {
	const INPUT string = ``
	const EXPECTED_PART_2 int = 0
	if resultPart2 := SolvePart2(INPUT); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %d not equal to expected %d", resultPart2, EXPECTED_PART_2)
	}
}

func TestInputPart1(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED_PART_1 int = 0
	if resultPart1 := SolvePart1(string(input)); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", resultPart1, EXPECTED_PART_1)
	}
}

func TestInputPart2(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED_PART_2 int = 0
	if resultPart2 := SolvePart2(string(input)); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %d not equal to expected %d", resultPart2, EXPECTED_PART_2)
	}
}
