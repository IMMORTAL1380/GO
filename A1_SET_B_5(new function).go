/* /*ASSIGNMENT 1
SET B
Q.5. WAP in go language to explain new function*/
package main

import "fmt"

type Sum struct {
	x_val int
	y_val int
}

func main() {

	/* Use the new function to perform allocation,
	which will return a pointer to the value's address.*/

	var p1 = new(int)
	fmt.Println(p1)
	var p = new(Sum)
	p.xval=10
	p.yval =20
	fmt.Println(p)
}
