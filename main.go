package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'done'")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task to add.")
			return
		}
		task := args[2]

		f, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(task + "\n"); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Added task:", task)
		}

	case "list":
		data, err := os.ReadFile("tasks.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		tasks := strings.Split(string(data), "\n")

		for i, task := range tasks {
			if task != "" {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
