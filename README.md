# TUI Multi-Window App

A simple multi-window terminal application built with Go and the Bubble Tea library.

## ASCII Diagram

```
╭───────────────────────────────────────────────────────────────────────────╮
│                           TUI App Navigation                              │
├───────────────────────────────────────────────────────────────────────────┤
│  ╭─────────╮   ╭─────────╮   ╭───────────╮   ╭───────────╮                │
│  │  List   │◀──▶│  Info   │◀──▶│  Results  │◀──▶│ Exercises │                │
│  ╰─────────╯   ╰─────────╯   ╰───────────╯   ╰───────────╯                │
│                                                                           │
│  Current View: Exercises (Grammar: Present Tenses)                        │
│  ┌───────────────────────────────────────────────────────┐                │
│  │ 1. I ___ (to be) a student.                           │                │
│  │    Your answer: [am] ✓                                │                │
│  │ 2. She ___ (to have) a cat.                           │                │
│  │    Your answer: [have] ~                              │                │
│  └───────────────────────────────────────────────────────┘                │
╰───────────────────────────────────────────────────────────────────────────╯
```

## Description

This application demonstrates a basic multi-page layout in a terminal user interface. You can navigate between four pages: "List", "Info", "Results", and "Exercises" using the `left arrow` and `right arrow` keys. The "List" page contains a simple checklist, the "Info" page displays a static message, the "Results" page shows the items selected from the checklist, and the "Exercises" page presents interactive learning exercises from a YAML file.

## Installation and Usage

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/kiquetal/tui-app
    ```
2.  **Navigate to the project directory:**
    ```bash
    cd tui-app
    ```
3.  **Run the application:**
    ```bash
    go run main.go
    ```

## Controls

*   `up`/`k`: Move the cursor up in the checklist (on the List page) or move between exercise input fields (on the Exercises page).
*   `down`/`j`: Move the cursor down in the checklist (on the List page) or move between exercise input fields (on the Exercises page).
*   `enter`/`space`: Toggle an item in the checklist (on the List page).
*   `enter`: On the Exercises page, press Enter to submit answers and view results.
*   `right arrow`: Navigate to the next page.
*   `left arrow`: Navigate to the previous page.
*   `q`/`ctrl+c`: Quit the application.
