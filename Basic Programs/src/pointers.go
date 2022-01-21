package main

import "fmt"

// Program 1:

// func main() {
// 	a := 20
// 	fmt.Println("Address:", &a)
// 	fmt.Println("Value:", a)
// }

// Program 2:

// func zero(x int) {
// 	x = 0
// }

// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
// }

// Program 3:

// func zero(x int) int {
// 	x = 0
// 	fmt.Println("Address inside other function:", &x)
// 	fmt.Println("Value inside other function:", x)
// 	return x
// }

// func main() {
// 	x := 5
// 	fmt.Println("Address before passing:", &x)
// 	fmt.Println("Value before passing:", x)
// 	x = zero(x)
// 	fmt.Println("Address after passing:", &x)
// 	fmt.Println("Value after passing:", x)
// 	fmt.Println(x) // x is still 5
// }

// Program 4:

func zero(xPtr *int) {
	fmt.Println("Address before passing:", &xPtr)
	fmt.Println("Pointer before passing:", *xPtr)
	fmt.Println("Value before passing:", xPtr)
	*xPtr = 0
	fmt.Println("Address after passing:", &xPtr)
	fmt.Println("Pointer after passing:", *xPtr)
	fmt.Println("Value after passing:", xPtr)
}

func main() {
	x := 5
	fmt.Println("Address before passing:", &x)
	fmt.Println("Value before passing:", x)
	zero(&x)
	fmt.Println("Address after passing:", &x)
	fmt.Println("Value after passing:", x)
	fmt.Println(x) // x is 0
}

// Program 5:

// func main() {
// 	// Create an integer variable a with value 20
// 	a := 20
// 	// Create a pointer variable b and assigned the address of a
// 	var b *int = &a
// 	var c int = a
// 	// print address of a(&a) and value of a
// 	fmt.Println("Address of a:", &a)
// 	fmt.Println("Value of a:", a)
// 	// print b which contains the memory address of a i.e. &a
// 	fmt.Println("Address of pointer b:", b)
// 	// *b prints the value in memory address which b contains i.e. the value of a
// 	fmt.Println("Value of pointer b:", *b, "and Value of b:", b)
// 	// print c which contains the value a and address of c
// 	fmt.Println("Address of c:", &c)
// 	fmt.Println("Value of c:", c)
// 	//increment the value of variable a using the variable b
// 	*b = *b + 1
// 	//prints the new value using a and *b
// 	fmt.Println("Value of pointer after opoeration on b:", *b)
// 	fmt.Println("Value of a:", a)
// }

// Program 6:

// func one(xPtr *int) {
// 	fmt.Println("Inside Function Before value of xPtr:", *xPtr, ", Pointer Value of xPtr", xPtr, "and address of xPtr:", &xPtr)
// 	*xPtr = 1
// 	fmt.Println("Inside Function After value of xPtr:", *xPtr, ", Pointer Value of xPtr", xPtr, "and address of xPtr:", &xPtr)
// }

// func main() {
// 	xPtr := new(int)
// 	fmt.Println("Before value of xPtr:", *xPtr, ", Pointer Value of xPtr", xPtr, "and address of xPtr:", &xPtr)
// 	one(xPtr)
// 	fmt.Println("After value of xPtr:", *xPtr, ", Pointer Value of xPtr", xPtr, "and address of xPtr:", &xPtr)
// 	fmt.Println(*xPtr) // x is 1
// }
