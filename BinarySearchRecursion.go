package main

import "fmt"

func BinarySearchRecursive(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	if BinarySearchRecursive(items, 0, len(items)-1, 31) {
		fmt.Println("Item is in the list")
	} else {
		fmt.Println("Item is not in the list")
	}
}
