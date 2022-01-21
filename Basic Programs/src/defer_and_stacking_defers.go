package main

import "fmt"

// Program 1(Stacking Defers):

// func main() {
// 	defer display(1)
// 	defer display(2)
// 	defer display(3)
// 	fmt.Println(4)
// }

// func display(a int) {
// 	fmt.Println(a)
// }

// Program 2(Defer):

func sample() {
	fmt.Println("Inside the sample()")
}
func main() {
	//sample() will be invoked only after executing the statements of main()
	defer sample()
	fmt.Println("Inside the main()")
}
