package main

import (
	"fmt"
	"strconv" //to convert int to string
)

func main() {
	var i int //declaring the variable : var name type
	i = 42    //initializing the variable
	fmt.Println(i)

	var j int = 27 // declaring and initializing
	fmt.Println(j)

	k := 10.0 //automatically define the right type
	fmt.Printf("%v, %T\n", k, k)

	var m float64
	m = float64(i) //type conversion
	fmt.Printf("%v, %T\n", m, m)

	var n string
	n = string(i) //converts to equivalent ASCII codes
	fmt.Printf("%v, %T\n", n, n)

	n = strconv.Itoa(j) //converts to string
	fmt.Printf("%v, %T\n", n, n)
}
