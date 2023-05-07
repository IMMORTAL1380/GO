/*
ASSIGNMENT 3
SET C
1. WAP in go language to accept two matrices and display it's
 multiplication */
// Golang program to multiply two matrices.

package main

import "fmt"

func main() {
	var sum int = 0
	var matrix1 [2][2]int
	var matrix2 [2][2]int
	var matrix3 [2][2]int

	fmt.Printf("Enter matrix1 elements: \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Elements: matrix1[%d][%d]: ", i, j)
			fmt.Scanln(&matrix1[i][j])
		}
	}

	fmt.Printf("Enter matrix2 elements: \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Elements: matrix2[%d][%d]: ", i, j)
			fmt.Scanln(&matrix2[i][j])
		}
	}

	//Multiplication of matrix1 and matrix2.
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum = 0
			for k := 0; k < 2; k++ {
				sum = sum + matrix1[i][k]*matrix2[k][j]
			}
			matrix3[i][j] = sum
		}
	}

	fmt.Printf("Matrix1: \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Matrix2: \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix2[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Multiplication of matrix1 and matrix2: \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix3[i][j])
		}
		fmt.Printf("\n")
	}
}
