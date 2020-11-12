package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "Welcome to KloudOne!" }()

	msg := <-messages
	fmt.Println(msg)
}
