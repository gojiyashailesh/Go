package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tasks []string

func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("✅ Task added")
}

func viewTask() {
	if len(tasks) == 0 {
		fmt.Println("📝 No tasks exist")
		return
	}
	fmt.Println("📋 Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func removeTask(index int) error {
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("❌ Invalid index")
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("🗑️ Task removed")
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n -----> TODO MANAGER <-------")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. View Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose the option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			addTask(task)

		case "2":
			viewTask()
			fmt.Print("Enter the task number to remove: ")
			scanner.Scan()
			numStr := scanner.Text()
			index, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Please enter a valid number")
				continue
			}
			err = removeTask(index - 1)
			if err != nil {
				fmt.Println(err)
			}

		case "3":
			viewTask()

		case "4":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid option. Try again.")
		}
	}
}
