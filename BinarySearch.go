package main

import "fmt"

func binarySearch(arr []int, val int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < val {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != val {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	if binarySearch(items, 56) {
		fmt.Println("Item is in the list")
	} else {
		fmt.Println("Item is not in the list")
	}
}
