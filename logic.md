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

## 2. Evaluating Answers

Once the exercises are loaded, the application will present the questions to the user one by one. For each question, the user will provide an answer. The evaluation process is as follows:

1.  **Store user answers:** As the user answers each question, their answers should be stored in a data structure, for example, a map or a slice.

2.  **Compare user answers with correct answers:** After the user has answered all the questions, the application will compare the user's answers with the correct answers from the `ExerciseData` struct. The comparison should be case-insensitive to avoid penalizing the user for capitalization differences.

    ```go
    import "strings"

    func checkAnswer(userAnswer, correctAnswer string) bool {
        return strings.ToLower(userAnswer) == strings.ToLower(correctAnswer)
    }
    ```

## 3. Generating a Final Report

After evaluating all the answers, the application will generate a final report. The report should provide the user with feedback on their performance. The report should include:

1.  **Score:** The number of correct answers and the total number of questions (e.g., "You got 4 out of 6 correct").

2.  **Detailed feedback:** For each question, the report should show the user's answer and the correct answer, especially for the questions they got wrong.

    **Example Report:**

    ```
    Your score: 4/6

    ---

    ### Detailed Report:

    1. I ___ (to be) a student.
       - Your answer: am (Correct)

    2. She ___ (to have) a cat.
       - Your answer: have (Incorrect)
       - Correct answer: has

    3. They ___ (to go) to school every day.
       - Your answer: go (Correct)

    4. He ___ (to play) football on weekends.
       - Your answer: play (Incorrect)
       - Correct answer: plays

    5. We ___ (to study) English.
       - Your answer: study (Correct)

    6. The sun ___ (to rise) in the east.
       - Your answer: rises (Correct)
    ```
