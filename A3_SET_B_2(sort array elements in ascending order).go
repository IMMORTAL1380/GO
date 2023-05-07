/*
ASSIGNMENT 3
SET B
Q.2. WAP in go language to sort array elements in ascending order.
*/
// Golang program to sort an integer array in
// ascending order using selection sort

package main

import "fmt"

func main() {
	var arr [5]int
	var min int = 0
	var temp int = 0

	fmt.Printf("Enter 5 array elements: \n")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elements: arr[%d]: ", i)
		fmt.Scanln(&arr[i])
	}

	for i := 0; i <= 4; i++ {
		min = i
		for j := i + 1; j <= 4; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}

	fmt.Printf("Array after sorting: \n")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
