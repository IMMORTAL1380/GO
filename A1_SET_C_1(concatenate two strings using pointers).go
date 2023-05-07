/*ASSIGNMENT 1
SET C
Q.1 WAP in go language to concatenate two strings using pointers.*/
package main

import (
	"fmt"
)

func main() {
	var string1 string
	var string2 string
	var str1 *string
	var str2 *string

	fmt.Println("Enter the first string")
	fmt.Scanln(&string1)
	fmt.Println("Enter the second string")
	fmt.Scanln(&string2)

	str1 = &string1
	str2 = &string2

	fmt.Println("string 1: ", *str1)
	fmt.Println("string 2: ", *str2)

	*str1 += *str2

	fmt.Printf("The concatenated string is %s", *str1)
}
