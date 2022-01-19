package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	//Example of different type of operator

	fmt.Printf("Line 1 - Type of variable a is = %T\n", a)
	fmt.Printf("Line 2 - Type of variable b is = %T\n", b)
	fmt.Printf("Line 3 - Type of variable c is = %T\n", c)

	//Example of & and * operators

	ptr = &a // now 'ptr' contains the address of 'a'
	fmt.Printf("value of a is %d\n", a)
	fmt.Printf("value of *ptr is %d\n", *ptr)
	fmt.Printf("value of ptr is %d\n", ptr)
	fmt.Printf("value of &a is %d\n", &a)
}
