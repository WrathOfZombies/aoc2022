package day5

import (
	"fmt"
	"regexp"
	"strings"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

type stack []string
type stackConfig []stack

type instruction struct {
	src  int // source stack
	dest int // destination stack
	len  int // number of items to move
}

// https://adventofcode.com/2022/day/5
// Given a stack and a list of instructions, determine the final arrangement of the stack
func SolvePart1(input string) string {
	cfg, instructions := parseInput(input)

	for _, instruction := range instructions {
		cfg = moveItem(cfg, instruction)
	}

	return top(cfg)
}

// Given a string input, return a stacks struct and a slice of instructions
func parseInput(input string) (stackConfig, []instruction) {
	// from the input we know that they are separated by a single line
	data := strings.Split(input, "\n\n")
	stacksStr, instructionStr := data[0], data[1]
	return parseStacks(stacksStr), parseInstructions(instructionStr)
}

// Given a stack configuration, parse it into a stacks struct
func parseStacks(stacksStr string) stackConfig {
	lines := strings.Split(stacksStr, "\n")
	columnIndexs := findColumnIndices(lines[len(lines)-1])
	config := make(stackConfig, len(columnIndexs))

	// start at the second to last line and work backwards
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]

		// only check the column indxes
		for cIdx, colIdx := range columnIndexs {
			// comparing against the rune value of the character
			if line[colIdx] >= 'A' && line[colIdx] <= 'Z' {
				container := config[cIdx]
				container = append(container, string(line[colIdx]))
				config[cIdx] = container
			}
		}
	}

	fmt.Printf("found configuration: %v\n", config)
	return config
}

// Given an instruction line config, figure out where the column data is going to appear
// the idea is to use these indexes to scan the data in a columnar fashion to find the stack data
// note: the output is zero indexed
func findColumnIndices(lineConfig string) []int {
	result := []int{}

	// regex to match every integer on the line
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(lineConfig, -1)

	// loop through the match and its location
	for _, i := range matches {
		result = append(result, i[0])
	}

	fmt.Printf("found column indices: %v\n", result)
	return result
}

func parseInstructions(instructionStr string) []instruction {
	// regex to parse integers from "move 2 from 1 to 3"
	result := []instruction{}
	lines := strings.Split(instructionStr, "\n")

	for _, line := range lines {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		result = append(result, instruction{
			src:  stringUtils.ParseInt(matches[1]) - 1,
			dest: stringUtils.ParseInt(matches[2]) - 1,
			len:  stringUtils.ParseInt(matches[0]),
		})
	}

	fmt.Printf("found instructions: %v\n", result)
	return result
}

func moveItem(cfg stackConfig, instruction instruction) stackConfig {
	srcStack := cfg[instruction.src]
	destStack := cfg[instruction.dest]

	for i := 0; i < instruction.len; i++ {
		item := srcStack[len(srcStack)-1]
		srcStack = srcStack[:len(srcStack)-1]
		destStack = append(destStack, item)
	}

	cfg[instruction.src] = srcStack
	cfg[instruction.dest] = destStack
	return cfg
}

func top(cfg stackConfig) string {
	result := ""
	for _, stack := range cfg {
		result += stack[len(stack)-1]
	}
	return result
}
