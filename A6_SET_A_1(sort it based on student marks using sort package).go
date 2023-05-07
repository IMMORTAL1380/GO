// write a go program to create student struct with student name and
//marks and sort it based on student marks using sort package

package main

import (
	"fmt"
	"sort"
)

type student struct {
	name  string
	marks int
}

func main() {
	// Create a slice of student structs
	students := []student{
		{name: "Amol", marks: 85},
		{name: "amey", marks: 92},
		{name: "shubh", marks: 78},
		{name: "rohan", marks: 91},
		{name: "harshal", marks: 87},
	}

	// Sort the slice based on student marks
	sort.Slice(students, func(i, j int) bool {
		return students[i].marks > students[j].marks
	})

	// Print the sorted slice
	fmt.Println("Sorted students:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.name, s.marks)
	}
}
