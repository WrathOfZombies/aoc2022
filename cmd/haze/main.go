package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wrathofzombies/aoc2022/internal/haze"
)

func main() {
	program := haze.New()
	if _, err := tea.NewProgram(program).Run(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
