package main

import (
	"fmt"
	"os"
)

func main() {
	a := 20
	// a := 30 //Error as its already declared
	fmt.Println(a)

	s := "Hello"
	fmt.Println("value of s before conditional check is", s)

	if s[1] != 'e' {
		os.Exit(1)
	}
	s = "GoodBye"
	fmt.Println("value of s after conditional check and before using pointer is", s)

	var p *string = &s
	*p = "ciao"
	fmt.Println("value of s after using pointer is", s)
}
