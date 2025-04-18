package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID int
	Name string
	Age int
	Grades map[string]int
}

type NewStudent struct {
	Name string
	Age int
}

var classroom []*Student

//view all students
func viewAllStudents() {
	if len(classroom) == 0 {
		fmt.Printf("\n❌ No Student Exists in Classroom ❌\n")
		return
	}

	for _, student := range classroom {
		fmt.Printf("\n______________\nStudent ID: %d\nStudent name: %s\nStudent age: %d\nMaths: %d\nEng: %d\nPhy: %d\nChem: %d\nBio: %d\n_____________\n", student.ID, student.Name, student.Age, student.Grades["Maths"], student.Grades["Eng"], student.Grades["Phy"], student.Grades["Chem"], student.Grades["Bio"])
	}
}

//add new student 
func addNewStudent(s NewStudent) {
	//checks if student Name already exists
	for _, student := range classroom {
		if student.Name == s.Name {
			fmt.Printf("\n❌ Student already exists ❌\n")
			return
		}
	}

	//if not build new student profile
	id := len(classroom) + 1
	defaultGrade := map[string]int{
		"Maths": 0,
		"Eng": 0,
		"Phy": 0,
		"Chem": 0,
		"Bio": 0,
	}

	classroom = append(classroom, &Student{id, s.Name, s.Age, defaultGrade})
	fmt.Println("\n✅ Student has been added successfully")
}

//remove existing student
func removeStudent(id int) {
	for i, student := range classroom {
		if student.ID == id {
			fmt.Printf("\nRemoving %s profile\n", student.Name)
			classroom = append(classroom[:i], classroom[i+1:]...)
			fmt.Println("\n✅ Student profile has been deleted")
			return
		}
	}

	fmt.Printf("\n❌ Student with this ID does not exist ❌\n")
}

//update student data
func updateStudentData(id, mathScore, engScore, phyScore, chemScore, bioScore int) {
	var existingStudent *Student

	for _, student := range classroom {
		if student.ID == id {
			existingStudent = student
		}
	}

	updatedGrades := map[string]int{
		"Maths": mathScore,
		"Eng": engScore,
		"Phy": phyScore,
		"Chem": chemScore,
		"Bio": bioScore,
	}

	existingStudent.Grades = updatedGrades
}


func main() {
	defer fmt.Println("\nGood Bye.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("\n*********************************\nSTUDENT GRADEBOOK MANAGER v.0.9\n*********************************\n")
		fmt.Printf("\n1. View All Students\n2. Add New Student\n3. Remove a Student\n4. Update a Student data\n5. Exit\n")
		fmt.Printf("Enter a command: ")
		scanner.Scan()
		prompt := scanner.Text()

		switch prompt {
		case "1":
			viewAllStudents()
			fmt.Printf("Enter (e) to return: ")
			scanner.Scan()
			if key := scanner.Text(); key == "E" {
				return
			}
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
				return
			}

			addNewStudent(NewStudent{name, age})

			fmt.Printf("Enter (e) to return: ")
			scanner.Scan()
			if key := scanner.Text(); key == "E" {
				return
			}
		case "3":
			fmt.Printf("\nEnter Student ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			removeStudent(id)

			fmt.Printf("Enter (e) to return: ")
			scanner.Scan()
			if key := scanner.Text(); key == "E" {
				return
			}
		case "4":
			fmt.Printf("\nEnter Student ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Printf("❌ Unsupported input format ❌\n")
				return
			}

			for _, student := range classroom {
				if student.ID != id {
					fmt.Printf("\n❌ Student with this ID does not exist ❌\n")
					return
				}
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

			updateStudentData(id, mathScore, engScore, phyScore, chemScore, bioScore)
			fmt.Println("\n✅ Student grades have been updated successfully")

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