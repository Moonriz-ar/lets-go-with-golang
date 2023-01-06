package main

import (
	"fmt"
	"os"
)

// assignment: student system

// 1- add student
// 2- edit student
// 3- show all the students

func showMenu() {
	fmt.Println("Welcome to student system.")
	fmt.Println("1. Add student.")
	fmt.Println("2. Edit student info.")
	fmt.Println("3. Show all student's info.")
	fmt.Println("4. Exit system.")
}

// retrieve user input regarding student info
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("Please enter student info according to the format")
	fmt.Print("Please write student id: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("Please write student name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Please write student class: ")
	fmt.Scanf("%s\n", &class)
	return newStudent(id, name, class) // instanciate new student with student constructor function
}

func main() {
	// instanciate new student manager
	sm := newStudentMgr()

	for {
		// 1- print options menu
		showMenu()
		// 2- wait for user input option
		var input int
		fmt.Print("Please enter the menu number of your operation: ")
		fmt.Scanf("%d\n", &input)
		fmt.Println("The option you selected is: ", input)
		// 3- execute the option selected by user
		switch input {
		case 1:
			// add student
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// edit student info
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// show all student's info
			sm.getAllStudents()
		case 4:
			// exit system
			os.Exit(0)
		}
	}

}
