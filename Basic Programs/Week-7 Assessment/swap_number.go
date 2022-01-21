package main

import "fmt"

func main() {
	x, y := 10, 20
	fmt.Println("Before swapping x:", x, "y:", y)
	x, y = y, x
	fmt.Println("Before swapping x:", x, "y:", y)
}
