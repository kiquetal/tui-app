# TUI Multi-Window App

A simple multi-window terminal application built with Go and the Bubble Tea library.

## ASCII Diagram

```
╭───────────────────────────────────────────────────────────╮
│                   TUI App Navigation                      │
├───────────────────────────────────────────────────────────┤
│  ╭─────────╮   ╭─────────╮   ╭───────────╮                │
│  │  List   │◀──▶│  Info   │◀──▶│  Results  │                │
│  ╰─────────╯   ╰─────────╯   ╰───────────╯                │
│                                                           │
│  Current View: List (Checklist)                           │
│  ┌─────────────────────────────────┐                      │
│  │ > [ ] Buy carrot                │                      │
│  │   [ ] Buy celery                │                      │
│  │   [ ] Buy kohlrabi              │                      │
│  └─────────────────────────────────┘                      │
╰───────────────────────────────────────────────────────────╯
```

## Description

This application demonstrates a basic multi-page layout in a terminal user interface. You can navigate between three pages: "List", "Info", and "Results" using the `left arrow` and `right arrow` keys. The "List" page contains a simple checklist, the "Info" page displays a static message, and the "Results" page shows the items selected from the checklist.

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

*   `up`/`k`: Move the cursor up in the checklist (on the List page).
*   `down`/`j`: Move the cursor down in the checklist (on the List page).
*   `enter`/`space`: Toggle an item in the checklist (on the List page).
*   `right arrow`: Navigate to the next page.
*   `left arrow`: Navigate to the previous page.
*   `q`/`ctrl+c`: Quit the application.
