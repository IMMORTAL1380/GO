/*ASSIGNMENT 3
SET A
2. WAP in go language to accept the book details such as
BookID, Title, Author, Price. Read and display the
details of n number of books.
*/
package main

import "fmt"

type Books struct {
	book_id int
	title   string
	author  string
	price   int
}

func main() {
	var Book1 [10]Books /* Declare Book1 of type Book */
	var n int
	/* Accept the book details */
	fmt.Println("How many books...?")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter details of Book[%d] \n,i+1")
		fmt.Println("Enter the Book ID")
		fmt.Scanln(&Book1[i].book_id)
		fmt.Println("Enter the Book Title")
		fmt.Scanln(&Book1[i].title)
		fmt.Println("Enter the Book author")
		fmt.Scanln(&Book1[i].author)
		fmt.Println("Enter the Book Price")
		fmt.Scanln(&Book1[i].price)
	}
	/* print Books info */
	for i := 0; i < n; i++ {
		fmt.Printf("Book %d book_id : %d\n", i+1, Book1[i].book_id)
		fmt.Printf("Book %d title : %s\n", i+1, Book1[i].title)
		fmt.Printf("Book %d author : %s\n", i+1, Book1[i].author)
		fmt.Printf("Book %d subject : %d\n", i+1, Book1[i].price)
	}
}
