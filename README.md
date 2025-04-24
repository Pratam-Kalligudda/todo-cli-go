# 📝 Todo CLI App in Go

A simple and fast command-line todo list application built with Go. Manage your tasks from the terminal efficiently — add, list, update, mark progress, search, and delete tasks with ease!

## 📦 Installation

1. **Clone the repository:**

```bash
git clone https://github.com/Pratam-Kalligudda/todo-cli-go.git
cd todo-cli-go
```

2. **Build the app:**

```bash
if on windows: go build -o todo.exe
else : go build -o todo
```

3. **Run it:**

```bash
./todo <command> [arguments]
```

# Commands

Command | Description
add `<task title>` | Add a new task
list | List all tasks
list --done | List all completed tasks
list --in-progress | List all in-progress tasks
list --not-done | List all tasks not marked as done
mark-done `<task id>` | Mark task as done
mark-in-progress `<task id>` | Mark task as in-progress
update `<task id>` `<title>` | Update the title of a task
search `<keyword>` | Search tasks by keyword
delete `<task id>` | Delete a task
help / --help | Show help message

**Example**

./todo add "Finish writing blog post"
./todo list
./todo list --done
./todo mark-done 2
./todo update 3 "Update project documentation"
./todo search blog
./todo delete 1

**Tech Stack**
Language: Go (Golang)

https://roadmap.sh/projects/task-tracker
