package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("Unsorted :\n", slice)
	fmt.Println("Sorted :\n", mergeSort(slice), "\n")
}

// Generates a slice of random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
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
