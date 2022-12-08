package day5_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day5"
)

// Tests the example input
func TestExamplePart1(t *testing.T) {
	const INPUT string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	const EXPECTED_PART_1 string = "CMZ"
	if resultPart1 := SolvePart1(INPUT); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %s not equal to expected %s", resultPart1, EXPECTED_PART_1)
	}
}

// Tests the input file
func TestInputPart1(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED_PART_1 string = "JCMHLVGMG"
	if resultPart1 := SolvePart1(string(input)); !reflect.DeepEqual(resultPart1, EXPECTED_PART_1) {
		t.Errorf("Testing part 1: Output %s not equal to expected %s", resultPart1, EXPECTED_PART_1)
	}
}
