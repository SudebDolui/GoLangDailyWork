package main

import "fmt"

// Program 1:

// func main() {
// 	panic("PANIC")
// 	str := recover()
// 	fmt.Println(str)
// }

// Program 2:

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
