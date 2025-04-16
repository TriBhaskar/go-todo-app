# Go Todo Application with Cobra

## Project Setup

1. Initialize Go module:

```bash
go mod init go-todo-app
```

2. Install Cobra CLI tool:

```bash
go install github.com/spf13/cobra-cli@latest
```

3. Initialize Cobra project:

```bash
cobra-cli init
```

### For more info related to cobra go through

https://github.com/spf13/cobra
https://cobra.dev/#install
https://pkg.go.dev/github.com/spf13/cobra#Command

A command-line todo application built with Go and Cobra CLI framework.

## Features

- Add tasks
- List tasks
- Mark tasks as complete
- Delete tasks
- Store tasks persistently

## Installation

```bash
go get github.com/yourusername/go-todo-app
```

## Usage

The application can be used in two modes:

### Interactive Mode

```bash
go-todo-app
```

This will start the interactive todo manager:

```
_____       _______        _
  / ____|     |__   __|      | |
 | |  __  ___    | | ___   __| | ___
 | | |_ |/ _ \   | |/ _ \ / _  |/ _ \
 | |__| | (_) |  | | (_) | (_| | (_) |
  \_____|\___/   |_|\___/ \__,_|\___/

  Interactive Todo Manager

--------------------------------
Available commands:
  add [title]      - Add a new todo
  list             - List all todos
  list -a          - List all todos including completed ones
  done [id]        - Mark a todo as done
  rm [id]          - Remove a todo
  help             - Show this help menu
  quit, exit       - Exit the program
```

## Commands

- `add` - Add a new task
- `list` - List all tasks
- `done` - Mark a task as complete
- `rm` - Delete a task

## Dependencies

- [Cobra](https://github.com/spf13/cobra)

## License

MIT License
