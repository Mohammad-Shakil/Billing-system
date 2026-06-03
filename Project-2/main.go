package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  int
}

func main() {
	books := []Book{
		{Title: "(Go basics)", Author: "King Shakil", Price: 9999},
		{Title: "(Rich dad Poor dad)", Author: "Joseph", Price: 7},
		{Title: "(Human)", Author: "Robiul", Price: 10},
	}
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i].Title, books[i].Author, books[i].Price)
	}
}
