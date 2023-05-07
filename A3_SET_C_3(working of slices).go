/*
ASSIGNMENT 3
SET C
3. WAP in go language to demonstrate working of slices
(like append, remove, copy etc.)*/

package main

import "fmt"

func main() {

	//1. Create a slice and Appending in a slice
	fmt.Println("Create and Appending in a slice\n")
	primeNumbers := []int{2, 3}

	// add elements 5, 7 to the slice
	primeNumbers = append(primeNumbers, 5, 7)
	fmt.Println("Prime Numbers:", primeNumbers)

	//2. Creating a slice from an array
	fmt.Println("Creating a slice from an array\n")
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	sliceNumbers := numbers[4:7] // creating slice from an array

	fmt.Println(sliceNumbers)

	// 3. create a slice using make()
	fmt.Println("Creating a slice using make \n")
	numbers2 := make([]int, 5, 7)

	// add elements to numbers
	numbers2[0] = 13
	numbers2[1] = 23
	numbers2[2] = 33

	fmt.Println("Numbers2:", numbers2)

	// 4. Adding elements of one slice to another
	fmt.Println("Adding elements of one slice to another\n")
	// create two slices
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers:", evenNumbers)
	// 5. Copy one slice to another
	fmt.Println("Copy one slice to another\n")
	primeNumbers1 := []int{2, 3, 5, 7} // create two slices
	numbers3 := []int{1, 2, 3}

	// copy elements of primeNumbers to numbers
	copy(numbers3, primeNumbers1)

	fmt.Println("Numbers3:", numbers3) // print numbers
	// 6. find the length of a slice
	fmt.Println("find the length of a slice\n")
	numbers4 := []int{1, 5, 8, 0, 3} // create a slice of numbers
	length := len(numbers4)          // find the length of the slice

	fmt.Println("Length:", length)
	//7. Looping through a slice
	fmt.Println("Looping through a slice\n")
	numbers5 := []int{2, 4, 6, 8, 10}

	// for loop that iterates through the slice
	for i := 0; i < len(numbers5); i++ {
		fmt.Println(numbers5[i])
	}
	weather := []string{"Rainy", "Sunny", "Cloudy"}

	//8. changing the element of a slice
	fmt.Println("changing the element of a slice\n")
	weather[2] = "Stromy"
	fmt.Println(weather)

	//Remove an element from a slice
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.

	fmt.Println(a) // [A B E D]
}
