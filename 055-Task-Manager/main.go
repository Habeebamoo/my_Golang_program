package main

// welcome to my world :) let's ~GO

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//defining the task type
type Task struct {
	ID int
	Name string
	CreatedAt time.Time
}

//default task
var allTasks []Task

//add task function
func addTask(title string) {
	id := len(allTasks) + 1
	task := Task{
		ID: id,
		Name: title,
		CreatedAt: time.Now(),
	}

	allTasks = append(allTasks, task)
	fmt.Println("Task added ✅")
}

//delete task function
func deleteTask(id int) {
	var filteredTasks []Task

	for _, task := range allTasks {
		if task.ID != id {
			filteredTasks = append(filteredTasks, task)
		}
	}

	allTasks = filteredTasks
	fmt.Println("Task deleted ✅")
}

//view task function
func viewTasks() {
	if len(allTasks) == 0 {
		fmt.Println("No Task Available.")
	} else {
		for _, task := range allTasks {
			fmt.Printf("\nTask Id: %d\nTask Item: %s\nCreated At: %v.\n_________\n", task.ID, task.Name, task.CreatedAt)
		}
	}
}

//clear task function 
func clearTasks() {
	for _, task := range allTasks {
		deleteTask(task.ID)
	}
	fmt.Println("All Tasks Remove ✅")
}

func main() {
	defer fmt.Println("Good Bye.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\n🔊 Welcome to my Task-Manager Program.\n************\n")
		fmt.Println("\n1. Add Task\n2. Delete Task\n3. View task\n4. Clear Task\n5. Exit")
		fmt.Print("Choose an Option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter a task: ")
			scanner.Scan()
			title := scanner.Text()
			addTask(title)

		case "2":
			fmt.Print("Enter the task ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Print("ID not found")
			} else {
				deleteTask(id)
			}

		case "3":
			viewTasks()

		case "4":
			clearTasks()

		case "5":
			return

		default:
			fmt.Print("Invalid option, Try again")
		}
	}
}