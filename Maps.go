package main

import "fmt"

func main() {
	nameNo := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println(nameNo)
	fmt.Println(nameNo["C"]) //retrieving value corressponding to key
	nameNo["D"] = 4          //adding element
	fmt.Println(nameNo)
	delete(nameNo, "C") //deleting elements
	fmt.Println(nameNo)
	pop, ok := nameNo["E"] //not sure if the key is present
	fmt.Println(pop, ok)
	fmt.Println(len(nameNo)) // no. of elements
}
