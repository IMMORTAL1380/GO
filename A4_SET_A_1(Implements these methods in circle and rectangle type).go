/*
ASSIGNMENT 4
SET A
1. Write a program in go language to create an interface shape
that includes area and perimeter. Implements these methods in circle
and rectangle type.*/
package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	length  float64
	breadth float64
}

// Implementing methods of the tank interface
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func main() {
	var s shape

	s = circle{10}
	fmt.Printf("Area of circle : %.2f\n", s.area())
	fmt.Printf("Perimeter of the circle :%.2f\n", s.perimeter())

	s = rectangle{10, 20}
	fmt.Printf("Area of rectangle : %.2f\n", s.area())
	fmt.Printf("Perimeter of the rectangle :%.2f\n", s.perimeter())

}
