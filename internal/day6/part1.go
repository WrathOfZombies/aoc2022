package day6

// https://adventofcode.com/2022/day/6
// Given a signal, detect the number of characters read before a marker is detected
func SolvePart1(input string) int {
	return findMarker(input, 4)
}

func findMarker(input string, size int) int {
	found := false
	l := len(input)

	for cursor := 0; cursor < l; cursor++ {
		// If the cursor exceeds 4 character limit
		if cursor > l-size {
			break
		}

		// read 4 characters from the input into the buffer
		buffer := input[cursor : cursor+size]
		found = validateBuffer(buffer)

		if found {
			return cursor + size
		}
	}

	return 0
}

// Checks if the buffer has any duplicate characters
func validateBuffer(buffer string) bool {
	cache := make(map[string]bool)

	for _, c := range buffer {
		if _, dupe := cache[string(c)]; dupe {
			return false
		}
		cache[string(c)] = false
	}

	return true
}
