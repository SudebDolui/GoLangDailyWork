package main

import (
	"fmt"
)

func main() {

	// we use var that declares one or more variables with or without type
	// type is a keywork
	//var name type
	//var name = value
	//shorthand syntax
	name := 20 // dynamic type of variable according to the value stored
	myName := "Sudeb"
	fmt.Println(name, myName)

	// define i
	var i int // statically mentioning data type

	i = 10 // set value for i
	fmt.Println(i)

	// we can also set value as follows
	var y = 5 // style coming from JavaScript
	fmt.Println(y)

	var msg = "Remote host found."
	fmt.Println(msg)

	var foo, bar int = 100, 200 // Both foo and bar of type integer
	fmt.Println(bar)
	var my_integer, my_string = 100, "Hello" // Here my_integer is of type integer and my_string is of type string
	fmt.Println(my_integer, my_string)

	// shorthand syntax

	vehicle := "Mercedes"
	age := 52

	// Bool true or false

	var is_job_failed = false

	// print it

	fmt.Printf("%d %d %s\n", i, y, msg)

	fmt.Println(foo)

	fmt.Println(age)

	fmt.Println(vehicle)

	fmt.Println(is_job_failed)
}
