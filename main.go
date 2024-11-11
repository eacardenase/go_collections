package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

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

	// courseRatings := map[string]float64{}
	courseRatings := make(floatMap, 3)

	// no need for memory reallocation
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["node"] = 4.8

	// requires memory reallocation
	courseRatings["angular"] = 4.7

	// fmt.Println(courseRatings)
	courseRatings.output()
}
