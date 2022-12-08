package day5

// https://adventofcode.com/2022/day/5
// Given a stack and a list of instructions, determine the final arrangement of the stack
func SolvePart2(input string) string {
	cfg, instructions := parseInput(input)

	for _, instruction := range instructions {
		cfg = moveItemInOrder(cfg, instruction)
	}

	return top(cfg)
}

func moveItemInOrder(cfg stackConfig, in instruction) stackConfig {
	srcStack := cfg[in.src]
	destStack := cfg[in.dest]

	// Determine the point where we split the array
	// It's the length of the stack minus the number of items to move
	items := []string{}
	pivot := len(srcStack) - in.len
	srcStack, items = partition(srcStack, pivot)
	destStack = append(destStack, items...)

	cfg[in.src] = srcStack
	cfg[in.dest] = destStack
	return cfg
}

// Partitions the array into two slices
func partition(arr []string, pivot int) ([]string, []string) {
	return arr[:pivot], arr[pivot:]
}
