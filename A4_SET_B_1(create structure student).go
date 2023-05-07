/*
ASSIGNMENT 4
SET B
1. Write a program in go language to create structure
 student.Write a method show() whose receiver is a pointer
 of struct student.*/

package main

import "fmt"

// Student structure
type student struct {
	name   string
	branch string
	rollno int
}

// Method with a receiver of student type
func (a *student) show(abranch string) {
	(*a).branch = abranch
}

func main() {

	// Initializing the values of the Student structure
	s := student{name: "Sona",
		branch: "CSE",
		rollno: 2,
	}
	fmt.Println("Student's name: ", s.name)
	fmt.Println("Branch Name(Before): ", s.branch)
	fmt.Println("Branch Name(Before): ", s.rollno)
	// Creating a pointer
	p := s

	// Calling the show method
	p.show("ECE")
	fmt.Println("Student's name: ", s.name)
	fmt.Println("Branch Name(After): ", s.branch)
}
