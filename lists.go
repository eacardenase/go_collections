package main

import "fmt"

func main() {
	var bookNames [4]string = [4]string{"The Martian", "Project Hail Mary", "The Three Body Problem"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	bookNames[3] = "Cibola Burn"

	fmt.Println(prices)
	fmt.Println(bookNames)

	fmt.Println(prices[2])
}
