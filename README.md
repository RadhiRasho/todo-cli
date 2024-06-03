# todo-cli
Quick Todo CLI using Go and SQLite

![image](https://github.com/RadhiRasho/todo-cli/assets/54078496/658553be-102a-4d92-a343-15f971ae3241)

## Description

This is a simple CLI application that allows you to manage your tasks. You can add a new task with Name, Description and Priority. You can list all tasks auto sorted by Priority. You can delete a task by ID or by Name. You can also list all tasks with a specific priority.

The tasks are stored in a SQLite database. The database file is created in the user's temp directory with the name `Todos.sqlite`. The database file is created when the application is run for the first time.

## Features
- Add a new task with Name, Description and Priority
- List all tasks auto sorted by Priority
- Delete a task by ID
- Delete a task by Name (All Matching Tasks will be deleted)
- Help Command
- List all tasks with a specific priority

## Requirements
- Go 1.16 or higher

## Installation
```bash
go install github.com/RadhiRasho/todo-cli@latest
```

## Usage
```bash
todo add "Task Name" "Task Description" # Add a new task
todo add "Task Name" "Task Description" "Task Priority" # Add a new task with priority
todo list # List all tasks auto sorted by Priority
todo delete "Task ID" # Delete a task by ID
todo delete "Task Name" # Delete as task by Name (All Matching Tasks will be deleted)
todo help # Show help
todo prioirty "Priority" # List all tasks with a specific priority
```

## Example
```bash
todo add "Task 1" "Task 1 Description" "High"
todo add "Task 2" "Task 2 Description" "Low"
todo add "Task 3" "Task 3 Description" "Medium"
todo list
todo delete "Task 1"
todo list
```
