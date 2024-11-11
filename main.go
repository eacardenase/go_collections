package main

import "fmt"

func main() {
	// userNames := []string{} // index out of range error
	userNames := make([]string, 2, 5)

	fmt.Println(len(userNames), cap(userNames)) // 2 5

	userNames[0] = "Keyla"
	userNames[1] = "Samantha"

	fmt.Println(len(userNames), cap(userNames)) // 2 5

	userNames = append(userNames, "Edwin")
	userNames = append(userNames, "Alexander")

	fmt.Println(userNames)

	fmt.Println(len(userNames), cap(userNames)) // 4 5
}
