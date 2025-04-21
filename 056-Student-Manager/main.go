package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"os/exec"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//defining my errors
var (
	ErrStudentAlreadyExist = errors.New("\n❌ Student with this name already exists ❌\n")
	ErrStudentDoesNotExist = errors.New("\n❌ Student with this ID does not exist ❌\n")
	ErrNoStudentExists = errors.New("\n❌ No Student Exists in Classroom ❌\n")
)

type Student struct {
	ID int `json:"id"`
	Name string	`json:"fullname"`
	Age int `json:"age"`
	Grades map[string]int `json:"grades"`
}

type NewStudent struct {
	Name string
	Age int
}

func getDatabase() []Student {
	classroomData, err := os.ReadFile("db/classroom.json")
	handleError(err)

	var classroom []Student
	err = json.Unmarshal(classroomData, &classroom)
	return classroom
}

func updateDatabase(classroom []Student) {
	classroomJson, err := json.MarshalIndent(classroom, "", " ")
	handleError(err)

	err = os.WriteFile("db/classroom.json", classroomJson, 0644)
	handleError(err)
}

//view all students
func viewAllStudents() error {
	classroom := getDatabase()

	if len(classroom) == 0 {
		return ErrNoStudentExists
	}

	for _, student := range classroom {
		fmt.Printf("\n______________\nStudent ID: %d\nStudent name: %s\nStudent age: %d\nMaths: %d\nEng: %d\nPhy: %d\nChem: %d\nBio: %d\n_____________\n", student.ID, student.Name, student.Age, student.Grades["maths"], student.Grades["eng"], student.Grades["phy"], student.Grades["chem"], student.Grades["bio"])
	}
	return nil
}

//add new student 
func addNewStudent(s NewStudent) error {
	//checks if database exists
	classroomData, err := os.ReadFile("db/classroom.json")
	if err != nil {
		if os.IsNotExist(err) {
			er := os.MkdirAll("db", 0755)
			handleError(er)
		}
	}

	var classroom []*Student
	err = json.Unmarshal(classroomData, &classroom)
	handleError(err)

	//checks if student Name already exists
	for _, student := range classroom {
		if student.Name == s.Name {
			return ErrStudentAlreadyExist
		}
	}

	//if not build new student profile
	id := len(classroom) + 1
	defaultGrade := map[string]int{
		"maths": 0,
		"eng": 0,
		"phy": 0,
		"chem": 0,
		"bio": 0,
	}

	classroom = append(classroom, &Student{id, s.Name, s.Age, defaultGrade})
	fmt.Println("\n✅ Student has been added successfully")
	
	//updates the database
	classroomJson, err := json.MarshalIndent(classroom, "", " ")
	handleError(err)
	err = os.WriteFile("db/classroom.json", classroomJson, 0644)
	handleError(err)

	return nil
}

//remove existing student
func removeStudent(id int) error {
	classroom := getDatabase()

	for i, student := range classroom {
		if student.ID == id {
			fmt.Printf("\nRemoving %s profile\n", student.Name)
			classroom = append(classroom[:i], classroom[i+1:]...)
			fmt.Println("\n✅ Student profile has been deleted")
			updateDatabase(classroom)
			return nil
		}
	}

	return ErrStudentDoesNotExist
}

//update student data
func updateStudentData(id, mathScore, engScore, phyScore, chemScore, bioScore int) error {
	classroomData, err := os.ReadFile("db/classroom.json")

	var classroom []*Student
	err = json.Unmarshal(classroomData, &classroom)
	handleError(err)

	//checks if student exist
	for _, student := range classroom {
		if student.ID == id {
			var existingStudent *Student

			for _, student := range classroom {
				if student.ID == id {
					existingStudent = student
				}
			}
		
			updatedGrades := map[string]int{
				"maths": mathScore,
				"eng": engScore,
				"phy": phyScore,
				"chem": chemScore,
				"bio": bioScore,
			}
		
			existingStudent.Grades = updatedGrades
		
			//updates the database
			classroomJson, err := json.MarshalIndent(classroom, "", " ")
			handleError(err)
			err = os.WriteFile("db/classroom.json", classroomJson, 0644)
			handleError(err)

			fmt.Println("\n✅ Student grades have been updated successfully")
			return nil
		}
	}

	return ErrStudentDoesNotExist
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
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
	defer fmt.Println("\nGood Bye.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearTerminal()
		fmt.Printf("\n*********************************\nSTUDENT GRADEBOOK MANAGER v.0.9\n*********************************\n")
		fmt.Printf("\n1. View All Students\n2. Add New Student\n3. Remove a Student\n4. Update a Student data\n5. Exit\n")
		fmt.Printf("Enter a command: ")
		scanner.Scan()
		prompt := scanner.Text()

		switch prompt {
		case "1":
			err := viewAllStudents()
			handleError(err)

			fmt.Printf("Enter any key to return: ")
			scanner.Scan()
		case "2":
			fmt.Printf("\nStudents Name (e.g john-doe): ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Printf("\nStudents Age: ")
			scanner.Scan()
			ageStr := scanner.Text()

			age, err := strconv.Atoi(strings.TrimSpace(ageStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
			} else {
				err := addNewStudent(NewStudent{name, age})
				handleError(err)
			}

			fmt.Printf("Enter any key to return: ")
			scanner.Scan()

		case "3":
			fmt.Printf("\nEnter Student ID: ")
			scanner.Scan()

			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
			} else {
				err := removeStudent(id)
				handleError(err)
			}

			fmt.Printf("Enter (e) to return: ")
			scanner.Scan()

		case "4":
			fmt.Printf("\nEnter Student ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			//maths score
			fmt.Printf("Enter Maths Score: ")
			scanner.Scan()
			mathScoreStr := scanner.Text()
			mathScore, err := strconv.Atoi(strings.TrimSpace(mathScoreStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			//eng score
			fmt.Printf("Enter English Score: ")
			scanner.Scan()
			engScoreStr := scanner.Text()
			engScore, err := strconv.Atoi(strings.TrimSpace(engScoreStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			//phy score
			fmt.Printf("Enter Physics Score: ")
			scanner.Scan()
			phyScoreStr := scanner.Text()
			phyScore, err := strconv.Atoi(strings.TrimSpace(phyScoreStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			//chem score
			fmt.Printf("Enter Chemistry Score: ")
			scanner.Scan()
			chemScoreStr := scanner.Text()
			chemScore, err := strconv.Atoi(strings.TrimSpace(chemScoreStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			//bio score
			fmt.Printf("Enter Biology Score: ")
			scanner.Scan()
			bioScoreStr := scanner.Text()
			bioScore, err := strconv.Atoi(strings.TrimSpace(bioScoreStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			err = updateStudentData(id, mathScore, engScore, phyScore, chemScore, bioScore)
			handleError(err)

			fmt.Printf("Enter (e) to return: ")
			scanner.Scan()
			if key := scanner.Text(); key == "E" {
				return
			}
		case "5":
			return
		default:
			fmt.Println("Try again")
		}
	}
}