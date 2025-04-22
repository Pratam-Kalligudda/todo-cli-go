package task

import (
	"fmt"
	"strings"
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

func DeleteTask(id int) error {
	tasks, err := ReadFromFile()

	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found to delete.")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err = WriteToFile(tasks); err != nil {
				return fmt.Errorf("error while writing file to tasks : %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no task with that id found : %w", err)

}

func MarkAsDone(id int) error {
	tasks, err := ReadFromFile()

	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found to mark as done.")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			if err = WriteToFile(tasks); err != nil {
				return fmt.Errorf("error while writing file to tasks : %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no task with that id found : %w", err)
}

func UpdateTitle(id int, title string) error {
	tasks, err := ReadFromFile()

	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found to update.")
		return nil
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			if err = WriteToFile(tasks); err != nil {
				return fmt.Errorf("error while writing file to tasks : %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no task with that id found : %w", err)

}

func ClearCompleted() error {
	tasks, err := ReadFromFile()

	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found to clear.")
		return nil
	}

	var updatedTask []Task
	for _, task := range tasks {
		if !task.Done {
			updatedTask = append(updatedTask, task)
		}
	}
	if err = WriteToFile(updatedTask); err != nil {
		return fmt.Errorf("error while writing file to tasks : %w", err)
	}
	return nil
}

func SearchByKeyword(keyword string) error {
	tasks, err := ReadFromFile()

	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found to search.")
		return nil
	}
	found := false
	for _, task := range tasks {
		if strings.Contains(task.Title, keyword) {
			fmt.Println(task)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("no task with that keyword found : %w", err)
	}
	return nil
}
