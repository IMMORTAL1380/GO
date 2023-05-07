/*
ASSIGNMENT 4
SET A
3. Write a program in go language to create structure author.
Write a method show() whose receiver is struct author.*/
package main

import (
	"fmt"
)

type author struct {
	name      string
	address   string
	particles int
	salary    float64
}

func (a author) show() {
	fmt.Printf("Author name: %s\n", a.name)
	fmt.Printf("Author Address: %s\n", a.address)
	fmt.Printf("Publlished articles: %d\n", a.particles)
	fmt.Printf("Author's salary: %.2f\n", a.salary)
}
func main() {
	v := author{name: "Ramesh",
		address:   "Pune",
		particles: 203,
		salary:    34000,
	}
	v.show()
}
