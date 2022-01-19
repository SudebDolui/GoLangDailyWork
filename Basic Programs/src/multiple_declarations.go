package main

import "fmt"

func main() {

	fmt.Println() //Declaring a empty line
	var x int
	x = 3              // Assiging value for x
	fmt.Print("x:", x) // Printing the value of x which is 3

	var y int = 20
	fmt.Print("y:", y) // Printing the value of y which is 20 on the same line

	var z = 50
	fmt.Print("z:", z) // Printing the value of z which is 50 on the same line

	var i, j = 100, "Hello"
	fmt.Println(i, j)
}
