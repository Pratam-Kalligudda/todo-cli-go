package main

import (
	"fmt"

	"github.com/Pratam-Kalligudda/todo-cli-go/task"
)

func main() {
	err := task.Add("Complete ListTaskFeature")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully added task")

	if err = task.ListAllTask(); err != nil {
		fmt.Println(err.Error())
		return
	}

}
