package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3} //declaring array and giving values
	fmt.Printf("Array1: %v\n", a)

	b := [...]int{1, 2, 3, 4} //unknown size
	fmt.Printf("Array2: %v\n", b)
	fmt.Printf("Array2 #2: %v\n", b[2])

	fmt.Printf("Length of Array2: %v\n", len(b)) //size of array

	//MATRIX
	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}
	matrix[2] = [3]int{0, 0, 1}
	fmt.Println(matrix)

	//SLICE
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //declaring and intializing slices
	fmt.Println(c)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c)) //gives capacity
	d := c[:]                            //slice of all elements
	e := c[3:]                           //slice from 4th element to end
	f := c[:6]                           //slice first 6 elements
	g := c[3:6]                          //slice the 4th, 5th, and 6th elements
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	c = append(c, 11) //adds 11 to slice
	fmt.Println(c)

	//make function
	m := make([]int, 3, 100) //making slices; type elements capacity
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))
}
