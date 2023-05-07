/*
ASSIGNMENT 6
SET C
2. WAP in Go language to print file information. 
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file information
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("Modified:", fileInfo.ModTime())
}
