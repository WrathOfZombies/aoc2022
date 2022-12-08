package day5_test

import (
	"os"
	"reflect"
	"testing"

	. "github.com/wrathofzombies/aoc2022/internal/day5"
)

// Tests the example input
func TestExamplePart2(t *testing.T) {
	const INPUT string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	const EXPECTED_PART_2 string = "MCD"
	if resultPart2 := SolvePart2(INPUT); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %s not equal to expected %s", resultPart2, EXPECTED_PART_2)
	}
}

// Tests the input file
func TestInputPart2(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Errorf("unable to read file: %v", err)
	}

	const EXPECTED_PART_2 string = "JCMHLVGMG"
	if resultPart2 := SolvePart2(string(input)); !reflect.DeepEqual(resultPart2, EXPECTED_PART_2) {
		t.Errorf("Testing part 2: Output %s not equal to expected %s", resultPart2, EXPECTED_PART_2)
	}
}
