# Logic for TUI Learning App

This document outlines the logic for importing exercises from a YAML file, evaluating user answers, and generating a final report.

## 1. Importing Exercises from YAML

The exercises are stored in a YAML file (e.g., `Process.md`). To import these exercises into the Go application, we need to follow these steps:

1.  **Read the YAML file:** The first step is to read the content of the YAML file into a byte slice. This can be done using the `os.ReadFile` function from the Go standard library.

2.  **Define a Go struct:** We need to define a Go struct that matches the structure of the YAML data. For our `Process.md` file, the struct would look like this:

    ```go
    type Exercise struct {
        Question string `yaml:"question"`
        Answer   string `yaml:"answer"`
    }

    type ExerciseData struct {
        Topic     string     `yaml:"topic"`
        SubTopic  string     `yaml:"sub_topic"`
        Sentences []Exercise `yaml:"sentences"`
    }
    ```

3.  **Parse the YAML:** We will use a YAML parsing library for Go, such as `gopkg.in/yaml.v3`, to unmarshal the YAML data into our Go struct.

    ```go
    import (
        "gopkg.in/yaml.v3"
        "os"
    )

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
    ```

## 2. Intelligent Evaluation

To provide a more user-friendly and intelligent evaluation, we can implement a more sophisticated evaluation logic.

### 2.1. Flexible Answer Matching

Instead of a strict string comparison, we can use a fuzzy string matching algorithm to tolerate minor typos. The **Levenshtein distance** is a good candidate for this. We can define a threshold for the distance, and if the user's answer is within that threshold, we can consider it correct or provide a hint.

```go
import "github.com/agext/levenshtein"

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
```

### 2.2. Providing Hints

If the user's answer is partially correct, the application can provide a hint. For example, if the correct answer is "plays" and the user enters "play", the app could say "Almost there! Check the verb conjugation."

### 2.3. Confidence Score

We can calculate a confidence score for each answer based on the Levenshtein distance. A lower distance means a higher confidence score.

-   **Distance 0:** Confidence 100% (Correct)
-   **Distance 1-2:** Confidence 75% (Partially correct, typo)
-   **Distance > 2:** Confidence 0% (Incorrect)

## 3. Advanced Reporting

The final report should be more than just a score. It should be a learning tool that helps the user understand their mistakes.

### 3.1. Categorization of Errors

We can try to categorize the errors to give the user more specific feedback. For example:

-   **Spelling Mistake:** If the answer is partially correct (Levenshtein distance <= 2).
-   **Grammar Mistake:** If the answer is a valid word but not the correct one (e.g., "play" instead of "plays"). This would require a more advanced NLP (Natural Language Processing) approach, but for simple cases, we can have a predefined list of common mistakes.

### 3.2. Detailed and Friendly Feedback

The report should be encouraging and provide clear explanations.

**Example Report:**

```
Great effort! Here's your report:

Your score: 4/6 (66%)

---

### Detailed Report:

1. I ___ (to be) a student.
   - Your answer: am (Correct)

2. She ___ (to have) a cat.
   - Your answer: have (Incorrect)
   - Correct answer: has
   - Tip: Remember to use 'has' for the third person singular (he, she, it).

3. They ___ (to go) to school every day.
   - Your answer: go (Correct)

4. He ___ (to play) football on weekends.
   - Your answer: play (Incorrect)
   - Correct answer: plays
   - Tip: For the present tense, we add an 's' to the verb for the third person singular.

5. We ___ (to study) English.
   - Your answer: study (Correct)

6. The sun ___ (to rise) in the east.
   - Your answer: rises (Correct)
```

### 3.3. Adaptive Learning (Future Improvement)

Based on the user's performance, the application could suggest the next topic to study or repeat the current one. For example, if the user makes many mistakes with verb conjugations, the app could suggest a lesson on that topic.