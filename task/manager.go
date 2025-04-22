package task

import (
	"fmt"
)

func Add(title string) error {
	tasks, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	task := Task{
		ID:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, task)
	err = WriteToFile(tasks)

	return err

}

func ListAllTask() error {
	tasks, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
	return nil
}
