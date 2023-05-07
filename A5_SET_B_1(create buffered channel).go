/*
SET B 
1. WAP in Go to create buffered channel, store few values in it and find channel capacity and length. Read values from channel and find modified length of a channel.
*/
//---------------------------------------------------------------
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))

	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))
}