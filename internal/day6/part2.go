package day6

// https://adventofcode.com/2022/day/6
// Given a signal, detect the number of characters read before a marker is detected
func SolvePart2(input string) int {
	return findMarker(input, 14)
}
