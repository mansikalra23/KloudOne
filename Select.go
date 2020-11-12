package main

import "fmt"

func fibonacci(c, end chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-end:
			fmt.Println("Bye.")
			return
		}
	}
}

func main() {
	c := make(chan int)
	end := make(chan int)
	var n int
	fmt.Print("Enter number : ")
	fmt.Scan(&n)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(<-c)
		}
		end <- 0
	}()
	fibonacci(c, end)
}
