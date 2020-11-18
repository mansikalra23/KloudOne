package main

import "fmt"

func transformList(data []int) []int {
	size := len(data)
	N := size / 2
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			data[N-i+2*j], data[N-i+2*j+1] = data[N-i+2*j+1], data[N-i+2*j]
		}
	}
	return data
}

func main() {
	items := []int{11, 12, 13, 21, 22, 23}
	fmt.Println("Transform of list is : ", transformList(items))

}
