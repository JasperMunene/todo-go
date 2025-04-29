package main

import (
	"fmt"
	"os"
	"strconv"
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

	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide the number of the task to delete.")
			return
		}

		// Parse the task number to delete
		idx, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task number:", args[2])
			return
		}

		// Read existing tasks
		data, err := os.ReadFile("tasks.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Split and filter out any empty lines
		raw := strings.Split(string(data), "\n")
		var tasks []string
		for _, t := range raw {
			if t != "" {
				tasks = append(tasks, t)
			}
		}

		// Check bounds
		if idx < 1 || idx > len(tasks) {
			fmt.Printf("Task #%d does not exist.\n", idx)
			return
		}

		// Remove the chosen task
		removed := tasks[idx-1]
		tasks = append(tasks[:idx-1], tasks[idx:]...)

		// Write the updated list back to the file
		out := strings.Join(tasks, "\n")
		if out != "" {
			out += "\n"
		}
		if err := os.WriteFile("tasks.txt", []byte(out), 0644); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Deleted task:", removed)
	default:
		fmt.Println("Unknown command:", command)
	}
}
