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
		err := task.ListAllTask()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: done <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = task.MarkAsDone(id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Marked as done.")
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

	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  add <task title>     - Add a new task")
	fmt.Println("  list                 - List all tasks")
	fmt.Println("  done <task id>       - Mark task as done")
	fmt.Println("  search <keyword>     - Search tasks by keyword")
}
