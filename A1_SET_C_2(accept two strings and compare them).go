/**ASSIGNMENT 1
SET C
Q.2. WAP in go language to accept two strings and compare them.*/
package main

import (
	"fmt"
)

func main() {
	str1 := "Computer"
	str2 := "Science"
	str3 := "computer"

	//String comparision using == operator

	if str1 == str2 {
		fmt.Printf("strings 1 and 2 are equal")
	} else {
		fmt.Println("strings 1 and 2 are equal")
	}
	if str1 != str3 {
		fmt.Println("String 1 and 3 are not equal")
	} else {
		fmt.Println("String 1 and 3 are equal")
	}
}
