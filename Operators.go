package main

import "fmt"

func main() {
	a := 10
	b := 3
	//INTEGERS
	//Arithmetic Operators
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)

	//Bitwise Operators
	fmt.Println("AND:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("NAND:", a&^b)

	//Shifting
	fmt.Println("Left Shift:", a<<b)
	fmt.Println("Right Shift:", a>>b)

	//STRING
	//Concatination
	s1 := "Chandler"
	s2 := "Bing"
	fmt.Printf("%v,%T", s1+s2, s1+s2)
}
