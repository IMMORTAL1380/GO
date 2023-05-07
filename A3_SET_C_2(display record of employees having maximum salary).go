/*
ASSIGNMENT 3
SET C
2. WAP in go language to accept n records of employee information
(eno,ename,salary) and display record of employees having
maximum salary.*/
package main

import "fmt"

// Defining a struct type
type employee struct {
	eno    int
	eName  string
	salary int
}

func main() {
	var emp [10]employee
	var i, n, max_salary int

	fmt.Println("How many employees...?")
	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		fmt.Printf("Enter Employee %d details\n", i+1)
		fmt.Println("Enter Employee Number\n")
		fmt.Scanln(&emp[i].eno)
		fmt.Println("Enter Employee Name\n")
		fmt.Scanln(&emp[i].eName)
		fmt.Printf("Enter Salary of Employee %d\n", i+1)
		fmt.Scanln(&emp[i].salary)
	}
	fmt.Println("Employee Details entered by you are....\n")
	for i = 0; i < n; i++ {
		fmt.Printf("Employee %d details\n", i+1)
		fmt.Println("Employee Number:\t", emp[i].eno)
		fmt.Printf("Employee Name:%s\t", emp[i].eName)
		fmt.Printf("Salary of employee:%d\n", emp[i].salary)
	}
	max_salary = emp[0].salary
	for i = 1; i < n; i++ {
		if emp[i].salary > max_salary {
			max_salary = emp[i].salary
		}
	}
	fmt.Printf("Max salary is %d", max_salary)
}
