package main

import "fmt"

func main() {
	var x = 50
	if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x equals 10")
	} else {
		fmt.Println("x is more than 10")
	}
}
