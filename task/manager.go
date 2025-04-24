package task

import (
	"fmt"
	"strings"
	"time"
)

func Add(title string) error {
	tasks, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	task := Task{
		ID:          len(tasks) + 1,
		Description: title,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	err = WriteToFile(tasks)

	return err

}

func List(status string, flag bool) error {
	tasks, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("error while loading tasks : %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	found := false
	for _, task := range tasks {
		checkStatus := false
		if flag {
			checkStatus = task.Status == status
		} else {
			checkStatus = task.Status != status
		}

		// fmt.Printf("Status : %v, checkStatus : %v, flag : %v\n", task.Status, checkStatus, flag)

		if checkStatus {
			fmt.Println(task)
			found = true
		} else if status == "" {
			fmt.Println(task)
			found = true
		}

	}
	if !found {
		return fmt.Errorf("nothing matches the status in your list")
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
	found := false
	for i, task := range tasks {
		if task.ID == id {
			found = true
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	if !found {
		return fmt.Errorf("no task with that id found : %w", err)
	}

	for i := range tasks {
		tasks[i].ID = i + 1
	}
	if err = WriteToFile(tasks); err != nil {
		return fmt.Errorf("error while writing file to tasks : %w", err)
	}
	return nil
}

func UpdateStatus(id int, status string) error {
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
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
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

	for i, _ := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = title
			tasks[i].UpdatedAt = time.Now()
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
		if task.Status != "done" {
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
		if strings.Contains(strings.ToLower(task.Description), strings.ToLower(keyword)) {
			fmt.Println(task)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("no task with that keyword found : %w", err)
	}
	return nil
}
