package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	test1()
}

func test1() {
	/* book 1 specification */
	var Book1 Books
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2 := Books{"Telecom Billing","Zara Ali","Telecom Billing Tutorial",6495700}
 	
 	printBook(Book1)
 	printBook_Pointer(&Book2) 
}

func printBook( book Books ){
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n\n", book.book_id);
}

func printBook_Pointer( book *Books ){
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n\n", book.book_id);
}