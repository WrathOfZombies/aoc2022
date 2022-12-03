package day1

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/wrathofzombies/aoc2022/internal/common/stringUtils"
)

// An array that implements the heap interface
type ElvesHeap []int

// Standard min heap implementation
// Copied over from godocs
// https://pkg.go.dev/container/heap
func (h ElvesHeap) Len() int           { return len(h) }
func (h ElvesHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h ElvesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ElvesHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *ElvesHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ElvesHeap) Insert(x any) {
	heap.Push(h, x)
	// If the size of the heap exceeds 3
	// then let's pop the smallest item out (hence the min heap)
	if h.Len() > 3 {
		heap.Pop(h)
	}
}

// https://adventofcode.com/2022/day/1
// Given an input of the calories carried by each elf,
// returns the top 3 elves carrying the most food
func SolvePart2(input string) int {
	line := strings.Split(input, "\n")
	minHeap := make(ElvesHeap, 0, 3)
	localMax := 0
	heap.Init(&minHeap)

	for _, item := range line {
		// If the line is blank, then we move on to the next elf
		if item == "" {
			minHeap.Insert(localMax)
			localMax = 0
			continue
		}

		// Otherwise parse the line to get the calorie count
		// and update our localmax
		calorie := stringUtils.ParseInt(item)
		localMax += calorie
	}

	// flush
	minHeap.Insert(localMax)
	localMax = 0

	fmt.Println(minHeap)

	total := 0
	for _, calories := range minHeap {
		total += calories
	}

	return total
}
