package day5_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day5"
)

// Tests the example input 
func TestExamplePart2(t *testing.T) {
	const INPUT string = ``
	const EXPECTED int = 0
	if result := SolvePart1(INPUT); !reflect.DeepEqual(result, EXPECTED) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", result, EXPECTED)
	}
}

// Tests the input file
func TestInputPart2(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED int = 0
	if result := SolvePart2(string(input)); !reflect.DeepEqual(result, EXPECTED) {
		t.Errorf("Testing part 1: Output %d not equal to expected %d", result, EXPECTED)
	}
}
