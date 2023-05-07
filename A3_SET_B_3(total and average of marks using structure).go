/*
ASSIGNMENT 3
SET B
3. WAP in go language to accept n student details like
 roll_no, stud_name, mark1,mark2, mark3. Calculate the
total and average of marks using structure.
*/
package main

import "fmt"

// Defining a struct type
type Student struct {
	rollno int
	Name   string
	mark1  int
	mark2  int
	mark3  int
}

func main() {
	var s [10]Student
	var i, n, total int
	var a1 float64

	fmt.Println("How many students...(Max 10) ?")
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		fmt.Println("Enter Student %d details", i+1)
		fmt.Println("Enter Roll Number")
		fmt.Scanln(&s[i].rollno)
		fmt.Println("Enter Student Name")
		fmt.Scanln(&s[i].Name)
		fmt.Println("Enter Marks of subject 1")
		fmt.Scanln(&s[i].mark1)
		fmt.Println("Enter Marks of subject 2")
		fmt.Scanln(&s[i].mark2)
		fmt.Println("Enter Marks of subject 3")
		fmt.Scanln(&s[i].mark3)
	}
	fmt.Println("Student Details entered by you are....")
	for i = 0; i < n; i++ {
		fmt.Printf("Student %d details\n", i+1)
		fmt.Println("Roll Number: \n", s[i].rollno)
		fmt.Printf("Name        :%s \n", s[i].Name)
		fmt.Printf("Subject 1 marks:%d\n", s[i].mark1)
		fmt.Printf("Subject 2 marks:%d \n", s[i].mark2)
		fmt.Printf("Subject 3 marks:%d \n", s[i].mark3)
		total = s[i].mark1 + s[i].mark2 + s[i].mark3
		a1 = float64(total) / 3
		fmt.Printf("Total=%d \t Average=%f", total, a1)
		fmt.Println()
	}
}
