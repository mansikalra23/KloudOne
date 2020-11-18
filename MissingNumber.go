package main

import (
	"fmt"
)

func missingNumber(data []int) int {
	var found int
	size := len(data)
	for i := 1; i <= size; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if data[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			fmt.Println(i, "number missing.")
			return 1
		}
	}
	fmt.Println("No Number Missing")
	return 0
}

func main() {
	items := []int{1, 2, 3, 5, 6, 7, 8, 9}
	missingNumber(items)
}
