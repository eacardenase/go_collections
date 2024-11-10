package main

import "fmt"

func main() {
	var bookNames [4]string = [4]string{"The Martian", "Project Hail Mary", "The Three Body Problem"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	bookNames[3] = "Cibola Burn"

	// featuredBooks := bookNames[:3]
	featuredBooks := bookNames[1:]
	highlightedPrices := featuredBooks[:1]

	fmt.Println(prices)
	fmt.Println(bookNames)

	fmt.Println(prices[2])
	// fmt.Println(prices[1:5]) // index out of bounds
	fmt.Println(prices[1:4])

	fmt.Println(featuredBooks)
	fmt.Println(highlightedPrices)
}
