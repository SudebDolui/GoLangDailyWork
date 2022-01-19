package main

import "fmt"

func main() {

	var x = 50
	if x < 10 {
		fmt.Println("x is less than 10")
	}
	fmt.Println("x is greater than 10")

	var n int
	fmt.Print("Enter your number: ")
	fmt.Scanf("%d", &n)
	fmt.Println(n)

	if n >= 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negative")
	}

	var y = 100
	if y < 10 {
		//Executes if y is less than 10
		fmt.Println("y is less than 10")
	} else if y >= 10 && y <= 90 {
		//Executes if x >= 10 and x<=90
		fmt.Println("y is between 10 and 90")
	} else {
		//Executes if both above cases fail i.e x>90
		fmt.Println("y is greater than 90")
	}

	fmt.Println(x == y)

	var a string = "hello"
	var b string = "hello"
	fmt.Println(a == b)
}
