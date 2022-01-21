package main

import (
	"calculation"
	"fmt"
)

func main() {
	x, y := 15, 10
	//the package will have function add()
	sum := calc.Add(x, y)
	fmt.Println("Sum:", sum)
	// fmt.Println(calc.Sum)
}
