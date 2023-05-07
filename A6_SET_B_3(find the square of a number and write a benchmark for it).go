/*
ASSIGNMENT 6
SET B
3. Write a function in Go language to find the square of a number and write a benchmark for it.
*/

package main

import (
	"fmt"
	"testing"
)

func Square(x int) int {
	return x * x
}

func main() {
	x := 5
	sq := Square(x)
	fmt.Printf("The square of %d is %d\n", x, sq)
}

// Benchmark for the Square function
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(5)
	}
}
