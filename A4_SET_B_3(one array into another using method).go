/*
ASSIGNMENT 4
SET B
3. Write a program in go language to copy all elements of one array
into another using method.*/
// Go program to illustrate how
// to copy an array by value

package main

import "fmt"

func main() {

      my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	my_arr2 := my_arr1 // Copy by value

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

	my_arr1[0] = "C++"

	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2: ", my_arr2)
}
//----------------------------------------------
// Copy by reference

// Go program to illustrate how to
// copy an array by reference

package main
import "fmt"
func main(){
my_arr1:= [6] int { 12, 456, 67, 65, 34, 34 }
my_arr2 : = &my_arr1 // Copy by refernce

fmt.Println("Array_1: ", my_arr1)
fmt.Println("Array_2:", *my_arr2)

my_arr1[5]= 1000000

fmt.Println("\nArray_1: ", my_arr1)
fmt.Println("Array_2:", *my_arr2)
}
