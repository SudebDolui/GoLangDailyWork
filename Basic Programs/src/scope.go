package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)
	f()
}

func f() {
	fmt.Println(x)
}

/*
Noticed that we moved the variable outside of the main function.
This means that other functions can access this variable:
*/
