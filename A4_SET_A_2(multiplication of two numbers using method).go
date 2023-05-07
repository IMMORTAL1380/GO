/*
ASSIGNMENT 4
SET A
2. Write a program in go language to print multiplication of two numbers using method*/

package main

import (
	"fmt"
)

type data int

func (num1 data) multiply(num2 data) data {
	return num1 * num2
}
func main() {
	var n1, n2 data
	fmt.Println("Enter first number")
	fmt.Scanln(&n1)
	fmt.Println("Enter second number")
	fmt.Scanln(&n2)
	result := n1.multiply(n2)
	fmt.Printf("The multiplication of %d and %d id %d", n1, n2, result)
}
