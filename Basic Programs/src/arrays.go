package main

import "fmt"

//Program 1:

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// }

//Program 2:

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83
// 	var total float64 = 0
// 	for i := 0; i < 5; i++ {
// 		total += x[i]
// 	}
// 	fmt.Println(total / 5)
// }

//Program 3:

// func main() {
// 	// define array named domains as string type
// 	var domains [2]string
// 	fmt.Println("current values for array:", domains)
// 	// add value and print it
// 	domains[0] = "cyberciti.biz"
// 	fmt.Println("Set value : ", domains)
// 	fmt.Println("Get value for 0 element : ", domains[0])
// 	// get array length
// 	fmt.Println("Array length : ", len(domains))
// 	// add one more value
// 	domains[1] = "nixcraft.com"
// 	// use for loop to print our array
// 	for i := 0; i < len(domains); i++ {
// 		fmt.Println("Get value for element ", i, " is ", domains[i])
// 	}
// }

//  Program 4:

func main() {
	var number [3]string // Decclaring a string array of size 3 and adding elements
	number[0] = "One"
	number[1] = "Two"
	number[2] = "Three"
	fmt.Println(number[1])                //prints Two
	fmt.Println(len(number))              //prints 3
	fmt.Println(number)                   // prints [One Two Three]
	directions := [...]int{1, 2, 3, 4, 5} // creating an integer array and the size of the array is defined by the number of elements
	fmt.Println(directions)               //prints [1 2 3 4 5]
	fmt.Println(len(directions))          //prints 5
	//Executing the below commented statement prints invalid array index 5 (out of bounds for 5-element array)
	// fmt.Println(directions[5])
}
