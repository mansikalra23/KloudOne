package main

import (
	"fmt"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		var piv = partition(arr, low, high)
		quickSort(arr, low, piv)
		quickSort(arr, piv+1, high)
	}
}

func partition(arr []int, low, high int) int {
	var piv = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= piv && i < high {
			i++
		}
		for arr[j] > piv && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = piv

	return j
}

func main() {

	var arr = []int{15, 3, 12, 6, -9, 9, 0}

	fmt.Println("Unsorted : ", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted : ", arr)
}
