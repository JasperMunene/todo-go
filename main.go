package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'done'")
		return
	}

	command := args[1]
	tasks := []string{"Walking the dog", "Cleaning", "Mowing the lawn"}

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task to add.")
			return
		}
		task := args[2]
		tasks = append(tasks, task)
		fmt.Println("Adding task:", task)

	case "list":
		for _, task := range tasks {
			fmt.Println(task)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
