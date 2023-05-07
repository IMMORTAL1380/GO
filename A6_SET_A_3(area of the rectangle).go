package main

import (
	"fmt"

	"rectangle" // Import the user-defined rectangle package
)

func main() {
	r := rectangle.Rectangle{Length: 4, Width: 6}

	area := rectangle.Area(r)

	fmt.Printf("The area of the rectangle is %.2f square units.\n", area)
}
