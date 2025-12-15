# Totion

Totion is a terminal-based note management application written in Go using Bubble Tea. It lets you create, edit, list, filter, and delete Markdown notes directly from the terminal. All notes are stored locally in a hidden directory inside the user’s home folder.

The project follows a clean internal package layout and separates TUI logic, filesystem access, styles, and utilities.

## Video Demo
<a href="https://drive.google.com/file/d/12aRVdFLY8uAXelexrkokOcCmnJCzvnFY/view?usp=sharing">Click Here</a><br>

## Features

* Create new Markdown notes
* Edit existing notes
* List and filter saved notes
* Delete notes with confirmation
* Fully keyboard-driven interface

## Storage

All notes(.md files) are stored in:

~/.totion

The directory is created automatically on first run.

## Project Structure

```
totion/
│
├── cmd/
│   └── totion/
│       └── main.go
│
├── internal/
│   ├── tui/
│   │   ├── model.go
│   │   ├── update.go
│   │   └── view.go
│   │
│   ├── fs/
│   │   └── files.go
│   │
│   ├── styles/
│   │   └── styles.go
│   │
│   └── util/
│       └── item.go
│
├── bin/
│   └── totion
│
├── .gitignore
├── Makefile
└── README.md
```

## Directory Overview

cmd/totion
Contains the application entry point

internal/tui
Handles Bubble Tea model, update, and view logic

internal/fs
Handles filesystem operations such as listing notes(.md files) and initializing the vault directory

internal/styles
Contains basic repeatative UI styling code

internal/util
Reusable helpers such as list item implementations

bin
Contains the compiled binary (not committed to version control)

## Key Bindings

Ctrl + N  Create a new note
Ctrl + L  List all notes
Enter     Open selected note or confirm an action
Ctrl + S  Save the current note
Ctrl + D  Delete the selected note
Esc       Go back or cancel the current action
Ctrl + C  Quit the application

## Build and Run

Requirements:
Go 1.20 or later

Build:

go build -o bin/totion ./cmd/totion

Run:

./bin/totion

## Dependencies

* <a href="github.com/charmbracelet/bubbletea">github.com/charmbracelet/bubbletea</a>
* <a href="github.com/charmbracelet/bubbles">github.com/charmbracelet/bubbles</a>
* <a href="github.com/charmbracelet/lipgloss">github.com/charmbracelet/lipgloss</a>

Dependencies are managed automatically using Go modules.

## Notes

* Notes are saved with a .md extension automatically
* Deleting a note requires typing the exact file name without the extension
* The application is fully keyboard-driven

## License

This project is open source and free to use.
