package main

// welcome to my world :) let's ~GO

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//defining the task type
type Task struct {
	ID 				 int 			`json:"id"`
	Name 			 string 		`json:"name"`
	Completed  bool			`json:"completed"`
	CreatedAt  time.Time	`json:"createdAt"`
}

//get tasks from json file
func getAllTasks() []*Task {
	allTaskData, err := os.ReadFile("db/tasks.json")
	handleError(err)

	var allTasks []*Task
	err = json.Unmarshal(allTaskData, &allTasks)
	handleError(err)

	return allTasks
}

//update tasks to json file
func updateAllTasks(allTasks []*Task) {
	allTaskJson, err := json.MarshalIndent(allTasks, "", " ")
	handleError(err)
	err = os.WriteFile("db/tasks.json", allTaskJson, 0644)
	handleError(err)
}

//add task function
func addTask(title string) error {
	//gets alltasks
	allTasksData, err := os.ReadFile("db/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			er := os.MkdirAll("db", 0755)
			handleError(er)
		}
	}

	//converts it
	var allTasks []Task
	err = json.Unmarshal(allTasksData, &allTasks)
	handleError(err)

	//checks if task exist
	for _, foundTask := range allTasks {
		if foundTask.Name == title {
			return errors.New("❌ Tasks already exists ❌")
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

	allTaskJson, err := json.MarshalIndent(allTasks, "", " ")
	handleError(err)
	err = os.WriteFile("db/tasks.json", allTaskJson, 0644)

	return nil
}

//delete task function
func deleteTask(id int) error {
	allTasks := getAllTasks()

	for i, task := range allTasks {
		if task.ID == id {
			allTasks = append(allTasks[:i], allTasks[i+1:]...)
			fmt.Println("Task deleted ✅")
			updateAllTasks(allTasks)
			return nil
		}
	}

	return errors.New("❌ Tasks ID does not exists ❌")
}

//view task function
func viewTasks() {
	db, err := os.ReadFile("db/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("❌ Task database does not exists ❌")
			return
		}
	}

	var allTasks []*Task
	err = json.Unmarshal(db, &allTasks)
	handleError(err)

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
	var allTasks []*Task
	allTaskJson, err := json.MarshalIndent(allTasks, "", " ")

	err = os.WriteFile("db/tasks.json", allTaskJson, 0644)
	handleError(err)
	fmt.Println("Task Cleared Succesfully ✅")
}

//complete task function
func completeTask(id int) error {
	allTasks := getAllTasks()

	for i := range allTasks {
		if allTasks[i].ID == id {
			allTasks[i].Completed = true
			fmt.Println("Task Completed ✅")
			updateAllTasks(allTasks)
			return nil
		}
	}

	return errors.New("❌ Tasks ID does not exists ❌")
}

//error handling function
func handleError(err error) {
	if err != nil {
		fmt.Printf("\n%s", err)
	}
}

//clear terminal
func clearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	defer fmt.Println("Good Bye.")
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		clearTerminal()
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
			err := addTask(title)
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
				err := deleteTask(id)
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
				err := completeTask(id)
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