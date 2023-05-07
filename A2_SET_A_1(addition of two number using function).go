//Assignment 2
//SET A 
//1. WAP in go language to print addition of two number using function.

package main
import "fmt"

func addNumber(x, y int) int {
   var sum int 
   sum =  x + y
   return sum
}
func main() {
   var number1, number2, number3 int
   number1 = 18
   number2 = 9
   number3 = addNumber(number1, number2)
   fmt.Println("The addition of ", number1," and ",number2,"is",number3)
}