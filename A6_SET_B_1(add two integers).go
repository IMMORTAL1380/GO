/*
Assignment 6
SET B 1. WAP in Go language to add two integers and write code for unit test to test this code. 
*/

package main

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	sum := Add(3, 5)
	fmt.Println("The sum of 3 and 5 is", sum)
}

// Unit test for the Add function
func TestAdd(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{2147483647, 1, -2147483648}, // Test integer overflow
	}
	for _, test := range tests {
		got := Add(test.x, test.y)
		if got != test.want {
			t.Errorf("Add(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}
