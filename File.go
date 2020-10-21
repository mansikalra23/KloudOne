package main

import (
	"fmt"
	"io/ioutil" //for file streaming
	"log"       //to log out of error
	"os"        //for file handling
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hey! Made by Mansi.")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)
	fmt.Println(s1)
}
