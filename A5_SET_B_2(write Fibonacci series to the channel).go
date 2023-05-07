/* Assignment 5 
SET B 
Q.2 */

//WAP in Go main go routine to read and write Fibonacci series to the channel
//-----------------------------------------------------------------
// The main function has two unbuffered channels ch and quit. 
//Inside fibonacci function the select statement blocks until 
//one of its cases is ready.
//------------------------------------------------------------------
package main

import (
	"fmt"
)

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x: // write to channel ch
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch) // read from channel ch
		}
		quit <- false
	}(n)

	fibonacci(ch, quit)
}