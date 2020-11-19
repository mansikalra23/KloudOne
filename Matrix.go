package main

import (
	"fmt"
)

func main() {
	var row, col int
	var mat [10][10]int
	fmt.Print("Enter no of rows: ")
	fmt.Scanln(&row)
	fmt.Print("Enter no of column: ")
	fmt.Scanln(&col)
	fmt.Println("\nEnter matrix elements: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&mat[i][j])
		}
	}
	fmt.Println("\nMatrix is: \n")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println("\n")
	}
}
