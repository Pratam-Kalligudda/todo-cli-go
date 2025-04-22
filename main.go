package main

import (
	"fmt"

	"github.com/Pratam-Kalligudda/todo-cli-go/task"
)

func main() {

	// if err := task.Add("Complete aws"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Successfully added task")

	if err := task.ClearCompleted(); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Successfully deleted task ")

	if err := task.ListAllTask(); err != nil {
		fmt.Println(err)
		return
	}

}
