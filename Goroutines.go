package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go one()
	go two()
	wg.Wait()
}
func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("one,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("two,  ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
