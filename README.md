# TUI Multi-Window App

A simple multi-window terminal application built with Go and the Bubble Tea library.

## ASCII Diagram

```
╭───────────────────────────────────────╮
│             TUI App Navigation        │
├───────────────────────────────────────┤
│  ╭───────────╮   ╭─────────╮          │
│  │ Lip Gloss │──▶│ Glamour │          │
│  ╰───────────╯   ╰─────────╯          │
│        ▲               │              │
│        │  (Tab Key)    │              │
│        └───────────────┘              │
│                                       │
│  Current View: Lip Gloss (Checklist)  │
│  ┌─────────────────────────────────┐  │
│  │ > [x] Buy carrot                │  │
│  │   [ ] Buy celery                │  │
│  │   [ ] Buy kohlrabi              │  │
│  └─────────────────────────────────┘  │
╰───────────────────────────────────────╯
```

## Description

This application demonstrates a basic multi-window layout in a terminal user interface. You can switch between the "Lip Gloss" and "Glamour" tabs using the `tab` key. The first tab contains a simple checklist, and the second tab displays a static message.

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

*   `up`/`k`: Move the cursor up in the checklist.
*   `down`/`j`: Move the cursor down in thechecklist.
*   `enter`/`space`: Toggle an item in the checklist.
*   `tab`: Switch between windows.
*   `q`/`ctrl+c`: Quit the application.
