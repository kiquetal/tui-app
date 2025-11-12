package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/agext/levenshtein"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#909090", Dark: "#909090"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	errorColor     = lipgloss.AdaptiveColor{Light: "#E88388", Dark: "#E88388"}

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

// Exercise represents a single question and its answer.
type Exercise struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

// ExerciseData represents the structure of the YAML file.
type ExerciseData struct {
	Topic     string     `yaml:"topic"`
	SubTopic  string     `yaml:"sub_topic"`
	Sentences []Exercise `yaml:"sentences"`
}

type model struct {
	currentPage   int
	choices       []string
	cursor        int
	selected      map[int]struct{}
	exercises     *ExerciseData
	userAnswers   []textinput.Model
	showResults   bool
	exerciseCursor int
}

// loadExercises reads and parses the exercises from a YAML file.
func loadExercises(filePath string) (*ExerciseData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var exerciseData ExerciseData
	err = yaml.Unmarshal(data, &exerciseData)
	if err != nil {
		return nil, err
	}

	return &exerciseData, nil
}

// checkAnswer evaluates the user's answer and returns if it's correct or partially correct.
func checkAnswer(userAnswer, correctAnswer string) (isCorrect bool, isPartial bool) {
	userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))
	correctAnswer = strings.ToLower(strings.TrimSpace(correctAnswer))

	if userAnswer == correctAnswer {
		return true, false
	}

	distance := levenshtein.Distance(userAnswer, correctAnswer, nil)
	if distance <= 2 { // Allow up to 2 typos
		return false, true // Partially correct
	}

	return false, false
}

func initialModel() model {
	exercises, err := loadExercises("Process.md")
	if err != nil {
		fmt.Printf("Error loading exercises: %v\n", err)
	}

	userAnswers := make([]textinput.Model, len(exercises.Sentences))
	for i := 0; i < len(exercises.Sentences); i++ {
		ti := textinput.New()
		ti.Placeholder = "Your answer"
		ti.Focus()
		userAnswers[i] = ti
	}

	return model{
		currentPage: 0,
		choices:     []string{"Buy carrot", "Buy celery", "Buy kohlrabi"},
		selected:    make(map[int]struct{}),
		exercises:   exercises,
		userAnswers: userAnswers,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("q", "ctrl+c"))):
			return m, tea.Quit
		case key.Matches(msg, key.NewBinding(key.WithKeys("right"))):
			if m.currentPage < 3 {
				m.currentPage++
			}
			if m.currentPage == 3 {
				m.userAnswers[m.exerciseCursor].Focus()
			} else {
				m.userAnswers[m.exerciseCursor].Blur()
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("left"))):
			if m.currentPage > 0 {
				m.currentPage--
			}
			if m.currentPage == 3 {
				m.userAnswers[m.exerciseCursor].Focus()
			} else {
				m.userAnswers[m.exerciseCursor].Blur()
			}
		}

		if m.currentPage == 0 {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
				if m.cursor > 0 {
					m.cursor--
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}

		if m.currentPage == 3 {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
				if m.exerciseCursor > 0 {
					m.userAnswers[m.exerciseCursor].Blur()
					m.exerciseCursor--
					m.userAnswers[m.exerciseCursor].Focus()
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
				if m.exerciseCursor < len(m.exercises.Sentences)-1 {
					m.userAnswers[m.exerciseCursor].Blur()
					m.exerciseCursor++
					m.userAnswers[m.exerciseCursor].Focus()
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
				m.showResults = true
			}
		}
	}

	if m.currentPage == 3 {
		m.userAnswers[m.exerciseCursor], cmd = m.userAnswers[m.exerciseCursor].Update(msg)
	}

	return m, cmd
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
	case 3:
		s.WriteString(fmt.Sprintf("Topic: %s\n", m.exercises.Topic))
		s.WriteString(fmt.Sprintf("Sub-Topic: %s\n\n", m.exercises.SubTopic))

		for i, exercise := range m.exercises.Sentences {
			s.WriteString(fmt.Sprintf("%d. %s\n", i+1, exercise.Question))
			s.WriteString(m.userAnswers[i].View())

			if m.showResults {
				isCorrect, isPartial := checkAnswer(m.userAnswers[i].Value(), exercise.Answer)
				if isCorrect {
					s.WriteString(lipgloss.NewStyle().Foreground(special).Render(" ✓"))
				} else if isPartial {
					s.WriteString(lipgloss.NewStyle().Foreground(errorColor).Render(" ~"))
				} else {
					s.WriteString(lipgloss.NewStyle().Foreground(errorColor).Render(" ✗"))
				}
			}
			s.WriteString("\n\n")
		}
		s.WriteString("\nPress Enter to see the results.\n")
	}

	// Page indicator
	pages := []string{"List", "Info", "Results", "Exercises"}
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