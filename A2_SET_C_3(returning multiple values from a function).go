/*
ASSIGNMENT 2
SET C
Q.3. WAP in go language to illustrate the concept of returning multiple values from a function
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(addAll(10, 15, 20, 25, 30))
}
func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0
	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}
