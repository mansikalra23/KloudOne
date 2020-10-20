package main

import "fmt"

type Friends struct {
	number int
	name   string
	friend []string
}

func main() {
	a := Friends{
		number: 3,
		name:   "Monica",
		friend: []string{
			"Chandler",
			"Rachel",
			"Joey",
		},
	}
	fmt.Println(a)
	fmt.Println(a.friend)

	b := a
	b.name = "Phoebe" //structs are independent datasets
	fmt.Println(b)
}
