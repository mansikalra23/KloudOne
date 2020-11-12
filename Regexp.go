package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("et$") // $ : the match must occur at the end of the string
	ra := regexp.MustCompile("..l*")
	r := regexp.MustCompile(".+l*")
	// . : wildcard character
	// *, +, {} : repeaters
	// returns the leftmost match

	fmt.Println(re.FindString("cricket"))
	fmt.Println(re.FindString("hacked"))
	fmt.Println(ra.FindString("rock and roll"))
	fmt.Println(r.FindString("rock and roll"))
}
