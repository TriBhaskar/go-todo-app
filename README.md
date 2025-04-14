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

```bash
# Add a new task
go-todo-app add "Buy groceries"

# List all tasks
go-todo-app list

# Mark task as complete
go-todo-app done 1

# Delete a task
go-todo-app delete 1
```

## Commands

- `add` - Add a new task
- `list` - List all tasks
- `done` - Mark a task as complete
- `delete` - Delete a task

## Dependencies

- [Cobra](https://github.com/spf13/cobra)

## License

MIT License
