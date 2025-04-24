// cmd/root.go
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Pratam-Kalligudda/todo-cli-go/task"
)

func Execute() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <task title>")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		err := task.Add(title)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task added successfully.")
		}

	case "list":
		if len(os.Args) < 3 {
			err := task.List("", true)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			property := os.Args[2]
			if property == "--done" {
				err := task.List("done", true)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else if property == "--in-progress" {
				err := task.List("in-progress", true)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else if property == "--not-done" {
				err := task.List("done", false)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-done <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = task.UpdateStatus(id, "done")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Marked as done.")
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-in-progress <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = task.UpdateStatus(id, "in-progress")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Marked as in-progress.")
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> <tilte>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		title := strings.Join(os.Args[3:], " ")
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = task.UpdateTitle(id, title)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Updated title.")
		}
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: search <keyword>")
			return
		}
		keyword := strings.Join(os.Args[2:], " ")
		err := task.SearchByKeyword(keyword)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <task id>")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err = task.DeleteTask(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Deleted title.")
		}
	case "--help", "help":
		printHelp()
	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  add <task title>            - Add a new task")
	fmt.Println("  list                        - List all tasks")
	fmt.Println("  list --done                 - List all completed tasks")
	fmt.Println("  list --in-progress          - List all in-progress tasks")
	fmt.Println("  list --not-done             - List all tasks not marked as done")
	fmt.Println("  mark-done <task id>         - Mark task as done")
	fmt.Println("  mark-in-progress <task id>  - Mark task as in-progress")
	fmt.Println("  update <task id> <title>    - Update the title of a task")
	fmt.Println("  search <keyword>            - Search tasks by keyword")
	fmt.Println("  delete <task id>            - Delete a task")
}
