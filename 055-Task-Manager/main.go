package main

// welcome to my world :) let's ~GO

import (
	"bufio"
	"errors"
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
	Completed bool
	CreatedAt time.Time
}

//default task
var allTasks []Task

//add task function
func addTask(title string) (int, error) {
	//checks if task exist
	for _, foundTask := range allTasks {
		if foundTask.Name == title {
			return 0, errors.New("❌ Tasks already exists ❌")
		}
	}

	var id int
	if len(allTasks) == 0 {
		id = 1
	} else {
		id = allTasks[len(allTasks) - 1].ID + 1
	}

	task := Task{
		ID: id,
		Name: title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	allTasks = append(allTasks, task)
	fmt.Printf("\nTask added ✅")

	return 1, nil
}

//delete task function
func deleteTask(id int) (int, error) {
	for i, task := range allTasks {
		if task.ID == id {
			allTasks = append(allTasks[:i], allTasks[i+1:]...)
			fmt.Println("Task deleted ✅")
			return 1, nil
		}
	}

	return 0, errors.New("❌ Tasks ID does not exists ❌")
}

//view task function
func viewTasks() {
	if len(allTasks) == 0 {
		fmt.Println("❌ No Task Available ❌")
	} else {
		for _, task := range allTasks {
			fmt.Printf("\nTask Id: %d\nTask Item: %s\nCompleted: %v.\nCreated At: %v.\n_________\n", task.ID, task.Name, task.Completed, task.CreatedAt)
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

//complete task function
func completeTask(id int) (int, error) {
	for i := range allTasks {
		if allTasks[i].ID == id {
			allTasks[i].Completed = true
			fmt.Println("Task Completed ✅")
			return 1, nil
		}
	}

	return 0, errors.New("❌ Tasks ID does not exists ❌")
}

//error handling function
func handleError(err error) {
	if err != nil {
		fmt.Printf("\n%s", err)
	}
}

func main() {
	defer fmt.Println("Good Bye.")
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Printf("\n🔊 Welcome to my Task-Manager Program.\n************\n")
		fmt.Printf("\n1. Add Task\n2. Delete Task\n3. View all tasks\n4. Clear Task\n5. Complete a Task\n6. Exit\n")
		fmt.Print("Choose an Option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter a task: ")
			scanner.Scan()
			title := scanner.Text()
			_, err := addTask(title)
			handleError(err)

			fmt.Printf("\nEnter any key to return: ")
			scanner.Scan()

		case "2":
			fmt.Print("Enter the task ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Print("ID not found")
			} else {
				_, err := deleteTask(id)
				handleError(err)
			}

			fmt.Printf("\nEnter any key to return: ")
			scanner.Scan()

		case "3":
			viewTasks()
			
			fmt.Printf("\nEnter any key to return: ")
			scanner.Scan()

		case "4":
			clearTasks()

			fmt.Printf("\nEnter any key to return: ")
			scanner.Scan()

		case "5":
			fmt.Print("Enter the task ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Print("❌ Invalid input format ❌")
			} else {
				_, err := completeTask(id)
				handleError(err)
			}

			fmt.Printf("\nEnter any key to return: ")
			scanner.Scan()

		case "6":
			return

		default:
			fmt.Print("Invalid option, Try again")
		}
	}
}