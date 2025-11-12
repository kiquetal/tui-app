package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#909090", Dark: "#909090"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	width, height int

	windowStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)
)

type model struct {
	currentPage int
	choices   []string
	cursor    int
	selected  map[int]struct{}
}

func initialModel() model {
	return model{
		currentPage: 0,
		choices:    []string{"Buy carrot", "Buy celery", "Buy kohlrabi"},
		selected:   make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("q", "ctrl+c"))):
			return m, tea.Quit
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
			// Toggle selection on the current item
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("right"))):
			if m.currentPage < 2 {
				m.currentPage++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("left"))):
			if m.currentPage > 0 {
				m.currentPage--
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	switch m.currentPage {
	case 0:
		s.WriteString("What should we buy at the market?\n\n")
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
		}
		s.WriteString("\nPress right arrow to go to the next page.\n")
	case 1:
		s.WriteString("This is the second window.\n\nPress left arrow to go back or right arrow to see the results.\n")
	case 2:
		s.WriteString("Here are your selected items:\n\n")
		if len(m.selected) == 0 {
			s.WriteString("You haven't selected any items.\n")
		} else {
			for i := range m.selected {
				s.WriteString(fmt.Sprintf("- %s\n", m.choices[i]))
			}
		}
		s.WriteString("\nPress left arrow to go back.\n")
	}

	// Page indicator
	pages := []string{"List", "Info", "Results"}
	pager := make([]string, len(pages))
	for i, pageName := range pages {
		if i == m.currentPage {
			pager[i] = lipgloss.NewStyle().Foreground(highlight).Underline(true).Render(pageName)
		} else {
			pager[i] = lipgloss.NewStyle().Foreground(subtle).Render(pageName)
		}
	}
	s.WriteString("\n\n" + strings.Join(pager, divider))

	s.WriteString("\nPress q to quit.\n")

	return windowStyle.Render(s.String())
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
