package haze

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	lavender     = lipgloss.Color("#EA9DFF")
	lavenderDark = lipgloss.Color("#39283F")
)

var (
	primaryStyle = lipgloss.NewStyle().Foreground(lavender)
	titleStyle   = lipgloss.NewStyle().
			Foreground(lavender).
			Background(lavenderDark).
			Padding(0, 1)
	cKey = key.NewBinding(
		key.WithKeys("c", "C"),
	)
)

type HazeModel struct {
	help         help.Model
	createMode   bool
	createPrompt CreatePromptModel
}

func New() HazeModel {
	return HazeModel{
		help:         help.New(),
		createMode:   false,
		createPrompt: CreatePrompt(),
	}
}

func (m HazeModel) Init() tea.Cmd {
	return nil
}

func (m HazeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil

		case key.Matches(msg, cKey):
			m.createMode = true
			m.createPrompt = CreatePrompt()
			return m, nil

		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	}

	// Keep trapping keystrokes until the user quits the create prompt
	if m.createMode {
		m.createPrompt, cmd = m.createPrompt.Update(msg)
		if m.createPrompt.Quit {
			m.createMode = false
		}
	}

	return m, cmd
}

func (m HazeModel) View() string {
	helpView := m.help.View(keys)
	height := 8 - strings.Count(helpView, "\n")
	msg := fmt.Sprintf(`Pick a problem from the list below. Press %s to create a new problem.`, primaryStyle.Render("'C'"))
	title := titleStyle.Render("WELCOME TO AOC 2022")

	if m.createMode {
		msg = m.createPrompt.View()
		title = titleStyle.Render("CREATE")
	}

	return fmt.Sprintf(`%s

%s
%s
%s
`, title, msg,
		strings.Repeat("\n", height),
		helpView)
}
