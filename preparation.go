package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/agext/levenshtein"
	"gopkg.in/yaml.v3"
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

func main() {
	// Load exercises from Process.md
	exerciseData, err := loadExercises("Process.md")
	if err != nil {
		fmt.Printf("Error loading exercises: %v\n", err)
		return
	}

	fmt.Printf("Topic: %s\n", exerciseData.Topic)
	fmt.Printf("Sub-Topic: %s\n\n", exerciseData.SubTopic)

	// Simulate user answers
	userAnswers := []string{"am", "have", "go", "play", "study", "rises"}

	var correctCount int
	var report strings.Builder

	report.WriteString("### Detailed Report:\n\n")

	for i, exercise := range exerciseData.Sentences {
		userAnswer := userAnswers[i]
		isCorrect, isPartial := checkAnswer(userAnswer, exercise.Answer)

		report.WriteString(fmt.Sprintf("%d. %s\n", i+1, exercise.Question))

		if isCorrect {
			correctCount++
			report.WriteString(fmt.Sprintf("   - Your answer: %s (Correct)\n\n", userAnswer))
		} else if isPartial {
			report.WriteString(fmt.Sprintf("   - Your answer: %s (Partially Correct)\n", userAnswer))
			report.WriteString(fmt.Sprintf("   - Correct answer: %s\n\n", exercise.Answer))
		} else {
			report.WriteString(fmt.Sprintf("   - Your answer: %s (Incorrect)\n", userAnswer))
			report.WriteString(fmt.Sprintf("   - Correct answer: %s\n\n", exercise.Answer))
		}
	}

	fmt.Printf("Your score: %d/%d\n\n", correctCount, len(exerciseData.Sentences))
	fmt.Println(report.String())
}
