/*ASSIGNMENT 1
SET C
Q.4 WAP in go language to check whether accepted number is
single digit or not.*/

package main

import (
	"fmt"
	"strconv"

)

func main() {
	var num int
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	if len(strings.Itoa(num) == 1 {
		fmt.Println("The number is a single digit number")
	}
}
