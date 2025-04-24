package task

import (
	"encoding/json"
	"fmt"
	"os"
)

var file string = "Tasks.json"

func WriteToFile(tasks []Task) error {

	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return fmt.Errorf("error while marshaling data: %w", err)
	}

	if err = os.WriteFile(file, data, 0664); err != nil {
		return fmt.Errorf("error while marshaling data : %w", err)
	}

	return nil

}

func ReadFromFile() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			// Create empty file if it doesn't exist
			emptyFile, createErr := os.Create(file)
			if createErr != nil {
				return nil, fmt.Errorf("failed to create file: %w", createErr)
			}
			emptyFile.Close()
			return []Task{}, nil
		}
		return nil, fmt.Errorf("error while reading file: %w", err)
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error while parsing data : %w", err)
	}

	return tasks, nil
}
