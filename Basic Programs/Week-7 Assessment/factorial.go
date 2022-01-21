package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	var user_input uint
	fmt.Print("Enter a String to check if vowel is present or not:")
	fmt.Scanf("%s\n", &user_input)
	println(factorial(user_input))
}
