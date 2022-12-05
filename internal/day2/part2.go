package day2

import (
	"fmt"
	"strings"
)

func toOutcome(input string) result {
	switch input {
	case "X":
		return loss
	case "Y":
		return draw
	case "Z":
		return win
	}
	return loss
}

func findCompliment(input choice, outcome result) choice {
	if outcome == draw {
		return input
	}

	if outcome == win {
		switch input {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		case empty:
			return empty
		}
	}

	switch input {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	case empty:
		return empty
	}

	return empty
}

// https://adventofcode.com/2022/day/2
// Given an input of the choices in the RockPaperScissors game, find the total score for the strategy
func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	totalScore := 0

	for _, line := range lines {
		fmt.Printf("\n\nInput: %s\n", line)

		// If the line is blank, then we move on to the next elf
		choices := strings.Split(line, " ")
		fmt.Printf("Choices: %v\n", choices)

		opponent := toChoice(choices[0])
		roundScore := computeScore(opponent, findCompliment(opponent, toOutcome(choices[1])))
		totalScore += roundScore
		fmt.Printf("TotalScore: %d, RoundScore: %d\n", totalScore, roundScore)
	}

	return totalScore
}
