/*
ASSIGNMENT 6
SET B
2. WAP in Go language to subtract two integers and write code for table test to test this code. 
*/

package main

import (
	"fmt"
	"testing"
)

func Subtract(x, y int) int {
	return x - y
}

func main() {
	diff := Subtract(5, 3)
	fmt.Println("The difference between 5 and 3 is", diff)
}

// Table test for the Subtract function
func TestSubtract(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{5, 3, 2},
		{3, 5, -2},
		{0, 0, 0},
		{-5, -3, -2},
		{2147483647, -1, -2147483648}, // Test integer overflow
	}
	for _, test := range tests {
		got := Subtract(test.x, test.y)
		if got != test.want {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}
