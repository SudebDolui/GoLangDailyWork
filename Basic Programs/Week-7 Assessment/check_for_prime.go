package main

import "fmt"

func CheckPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}

func main() {
	var user_input int
	fmt.Print("Enter a String to check if it is prime a number or not:")
	fmt.Scanf("%s\n", &user_input)
	CheckPrime(user_input)
}
