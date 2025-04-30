
# Simple CLI Task Manager

A basic command-line task manager written in Go. It allows users to add, list, and delete tasks stored in a local `tasks.txt` file.

## Features

- ✅ Add new tasks
- 📋 List all existing tasks
- ❌ Delete tasks by number

## Getting Started

### Requirements

- [Go](https://golang.org/dl/) 1.16 or higher

### Build

```bash
go build -o task
```

This will create an executable named `task` (or `task.exe` on Windows).

## Usage

### Add a Task

```bash
./task add "Buy groceries"
```

This appends the task to `tasks.txt`.

### List All Tasks

```bash
./task list
```

Example output:

```
1. Buy groceries
2. Read a book
```

### Delete a Task

```bash
./task delete 2
```

Deletes task number 2 (`Read a book` in the example above).

## Notes

- All tasks are saved in a file called `tasks.txt` in the current working directory.
- Deleted tasks are permanently removed from the file.
- Lines are trimmed for emptiness when listing or deleting.

## License

This project is licensed under the MIT License. Feel free to use and modify it for your own needs.
