# TODO App in Go

A simple command-line TODO list manager written in Go.

## Features

- Add new tasks.
- Complete tasks.
- Delete tasks.
- List all tasks.
- Update existing tasks.
- Store tasks in a local JSON file.

# Usage
- Add a new task
```go
./todo-app-go -add "Your task here"

```

- Complete a task

```go
./todo-app-go -complete 1

```
- Delete a task
```go
./todo-app-go -del 1

```

List all tasks:
```go
./todo-app-go -list

```

Update a task

```go
./todo-app-go -update 1 "Updated task here"

```