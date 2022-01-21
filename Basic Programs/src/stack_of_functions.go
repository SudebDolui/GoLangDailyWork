package main

import "fmt"

// Functions are built up in a “stack”.
// Suppose we had this program:

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}
