package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
   choices []string
   cursor int
   selected map[int]struct{}
}

func initialModel() model {

	return model{
		choices: []string{"Buy carrot","Buy celery","Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    return m, nil
}

func (m model) View() string {
    s := "What should we buy at the market?\n\n"

    for i, choice := range m.choices {

        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
		checked = "x" // selected!
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
		


