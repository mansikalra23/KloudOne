package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{23, 44, 45, 66, 67, 54, 48, 56, 88, 98, 19}
	if linearsearch(items, 56) {
		fmt.Println("Item is in the list")
	} else {
		fmt.Println("Item is not in the list")
	}
}
