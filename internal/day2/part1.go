package day2

import (
	"fmt"
	"strings"

	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

type choice int

const (
	empty    choice = 0
	rock            = 1
	paper           = 2
	scissors        = 3
)

type result int

const (
	loss result = 0
	draw        = 3
	win         = 6
)

// Returns the choice based on the input given by either player
func toChoice(input string) choice {
	switch input {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	}
	return empty
}

// Tells you whether you won the round
func getRoundOutcome(opponent choice, my choice) result {
	fmt.Printf("Opponent chose %v and I chose %v\n", opponent, my)

	if opponent == my {
		fmt.Println("We DREW.")
		return draw
	}

	if opponent == rock && my == paper {
		fmt.Println("I WON to a Rock with a Paper.")
		return win
	}

	if opponent == paper && my == scissors {
		fmt.Println("I WON to a Paper with a Scissors.")
		return win
	}

	if opponent == scissors && my == rock {
		fmt.Println("I WON to a Scissors with a Rock.")
		return win
	}

	fmt.Println("I LOST.")
	return loss
}

// Gives your total score for a round
func computeScore(opponent choice, my choice) int {
	outcome := getRoundOutcome(opponent, my)
	fmt.Printf("Outcome: %v, Score: %d\n", outcome, int(outcome))

	roundScore := int(outcome)
	roundScore += int(my)
	return roundScore
}

// https://adventofcode.com/2022/day/2
// Given an input of the choices in the RockPaperScissors game, find the total score for the strategy
func SolvePart1(input string) int {
	totalScore := 0

	for line := range stringUtils.GetLine(input) {
		fmt.Printf("\n\nInput: %s\n", line)

		// If the line is blank, then we move on to the next elf
		choices := strings.Split(line, " ")
		fmt.Printf("Choices: %v\n", choices)

		roundScore := computeScore(toChoice(choices[0]), toChoice(choices[1]))
		totalScore += roundScore
		fmt.Printf("TotalScore: %d, RoundScore: %d\n", totalScore, roundScore)
	}

	return totalScore
}
