package haze

import (
	"fmt"
	"os"
	"text/template"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	stringUtils "github.com/wrathofzombies/aoc2022/internal/common/string_utils"
)

var inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9e92f4"))

type CreatePromptModel struct {
	Input     textinput.Model
	Quit      bool
	Generated bool
}

type templateData struct {
	Day  string
	Part string
}

func CreatePrompt() CreatePromptModel {
	return CreatePromptModel{
		Input: inputModel(),
	}
}

func (m CreatePromptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreatePromptModel) Update(msg tea.Msg) (CreatePromptModel, tea.Cmd) {
	var cmd tea.Cmd

	// Handle keystrokes for this component
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Esc):
			m.Quit = true
			return m, nil

		case key.Matches(msg, keys.Enter):
			m.Quit = true
			m.createTemplate()
			if m.Generated {
				return m, nil
			}
		}
	}

	// Update the input component as well
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m CreatePromptModel) View() string {
	return fmt.Sprintf(
		"Pick a day from 1 to 25\n\n%s\n\n%s",
		m.Input.View(),
		"(esc to cancel)",
	) + "\n"
}

func (m CreatePromptModel) createTemplate() {
	if m.Input.Value() == "" {
		m.Generated = false
		return
	}

	input := stringUtils.ParseInt(m.Input.Value())
	if input > 25 || input < 1 {
		m.Generated = false
		return
	}

	// Working Directory
	wd, _ := os.Getwd()
	fmt.Printf("Working Directory: %s\n\n\n\n", wd)

	data := struct {
		Day string
	}{
		Day: fmt.Sprintf("%d", input),
	}

	outLocation := fmt.Sprintf("%s/internal/day%s/", wd, data.Day)

	info, _ := os.Lstat(outLocation)
	if info != nil {
		fmt.Printf("Directory %s already exists\n", outLocation)
		m.Generated = false
		return
	}

	os.Mkdir(outLocation, 0755)
	compileAndWrite("part.tmpl", outLocation+"part1.go", templateData{Day: data.Day, Part: "1"})
	compileAndWrite("part.tmpl", outLocation+"part2.go", templateData{Day: data.Day, Part: "2"})
	compileAndWrite("test.tmpl", outLocation+"part1_test.go", templateData{Day: data.Day, Part: "1"})
	compileAndWrite("test.tmpl", outLocation+"part2_test.go", templateData{Day: data.Day, Part: "2"})
	file, _ := os.Create(outLocation + "input")
	defer file.Close()

	m.Generated = true
}

func compileAndWrite(templateName string, outPath string, data templateData) {
	wd, _ := os.Getwd()
	templateLocation := fmt.Sprintf("%s/internal/haze/template/%s", wd, templateName)

	t := template.Must(template.New(templateName).ParseFiles(templateLocation))
	file, err := os.Create(outPath)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func inputModel() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter a day from 1 - 25"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return ti
}
