# Task CLI (Go)

A simple **command-line task manager** written in **Go**.

**Based on:** [Task Tracker (roadmap.sh)](https://roadmap.sh/projects/task-tracker)
It allows you to create, update, delete, and track tasks directly from the terminal.
Tasks are stored locally in a `tasks.json` file.

---

## Features

* Add new tasks
* Update existing tasks
* Delete tasks
* List all tasks
* Filter tasks by status
* Mark tasks as **in-progress**
* Mark tasks as **done**

---

## Requirements

* Go **1.18+** installed
* Terminal / command line

Check Go installation:

```bash
go version
```

---

## Installation

Clone the repository or copy the source file.

```bash
git clone <your-repository-url>
cd <project-folder>
```

Run the program:

```bash
go run todo-cli.go <command>
```

Optionally you can build a binary:

```bash
go build -o task-cli
```

Then run it:

```bash
./task-cli <command>
```

---

# Usage

## Add a Task

```bash
go run todo-cli.go add "Buy groceries"
```

Example:

```bash
go run todo-cli.go add "Finish Go project"
```

---

## Update a Task

```bash
go run todo-cli.go update <task_id> "New description"
```

Example:

```bash
go run todo-cli.go update 1 "Finish Go CLI project"
```

---

## Delete a Task

```bash
go run todo-cli.go delete <task_id>
```

Example:

```bash
go run todo-cli.go delete 2
```

---

## List All Tasks

```bash
go run todo-cli.go list
```

Example output:

```
[#1] [todo] Buy groceries
[#2] [in-progress] Finish project
[#3] [done] Read Go documentation
```

---

## List Tasks by Status

```bash
go run todo-cli.go list <status>
```

Available statuses:

* `todo`
* `in-progress`
* `done`

Examples:

```bash
go run todo-cli.go list todo
go run todo-cli.go list in-progress
go run todo-cli.go list done
```

---

## Mark Task as In Progress

```bash
go run todo-cli.go mark-in-progress <task_id>
```

Example:

```bash
go run todo-cli.go mark-in-progress 1
```

---

## Mark Task as Done

```bash
go run todo-cli.go mark-done <task_id>
```

Example:

```bash
go run todo-cli.go mark-done 1
```

---

# Task Structure

Tasks are stored in `tasks.json` with the following structure:

```json
{
  "id": 1,
  "description": "Finish project",
  "status": "todo",
  "created_at": "2026-03-05T12:00:00Z",
  "updated_at": "2026-03-05T12:00:00Z"
}
```

Fields:

| Field       | Description             |
| ----------- | ----------------------- |
| id          | Unique task identifier  |
| description | Task description        |
| status      | Current task status     |
| created_at  | Task creation timestamp |
| updated_at  | Last update timestamp   |

---

# Status Workflow

```
todo → in-progress → done
```

---

# Example Workflow

```bash
go run todo-cli.go add "Learn Go"
go run todo-cli.go add "Build CLI app"

go run todo-cli.go list

go run todo-cli.go mark-in-progress 1

go run todo-cli.go mark-done 1

go run todo-cli.go list done
```

---

# Data Storage

All tasks are saved in:

```
tasks.json
```

If the file does not exist, it will be created automatically when adding the first task.

---

# License

This project is free to use and modify.
