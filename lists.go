package main

import (
	"fmt"
)

// func main() {
// 	var bookNames [4]string = [4]string{"The Martian", "Project Hail Mary", "The Three Body Problem"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	bookNames[3] = "Cibola Burn"

// 	// featuredBooks := bookNames[:3]
// 	featuredBooks := bookNames[1:]

// 	featuredBooks[0] = "The Final Empire" // modifies the original array

// 	highlightedBooks := featuredBooks[:1]

// 	fmt.Println(prices)
// 	fmt.Println(bookNames)

// 	fmt.Println(prices[2])
// 	// fmt.Println(prices[1:5]) // index out of bounds
// 	fmt.Println(prices[1:4])

// 	fmt.Println(featuredBooks)
// 	fmt.Println(highlightedBooks) // [The Final Empire]
// 	fmt.Println(highlightedBooks[0])
// 	// fmt.Println(highlightedBooks[1]) // runtime error, although not at compile time

// 	fmt.Println("Book names", len(bookNames), cap(bookNames))                   // 4 4
// 	fmt.Println("Featured books", len(featuredBooks), cap(featuredBooks))       // 3 3
// 	fmt.Println("Featured books", len(highlightedBooks), cap(highlightedBooks)) // 1 3

// 	// highlightedBooks = highlightedBooks[0:] // runtime error
// 	highlightedBooks = highlightedBooks[:3]

// 	fmt.Println(highlightedBooks)                                               // [The Final Empire The Three Body Problem Cibola Burn]
// 	fmt.Println(highlightedBooks[1])                                            // The Three Body Problem
// 	fmt.Println("Featured books", len(highlightedBooks), cap(highlightedBooks)) // 3 3
// }

func main() {
	prices := []float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println("Prices", prices)
	fmt.Println(len(prices), cap(prices)) // 4 4

	// prices[4] = 4.99 // index out of range

	prices = append(prices, 5.99, 12.99)

	prices[4] = 4.99 // works, because a new array was created and replaced the old one

	fmt.Println("Prices", prices)
	fmt.Println(len(prices), cap(prices)) // 6 8
}
